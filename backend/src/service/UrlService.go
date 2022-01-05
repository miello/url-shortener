package service

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/miello/url-shortener/backend/src/dto"
	"github.com/miello/url-shortener/backend/src/types"
	"github.com/miello/url-shortener/backend/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateNewURL(ctx *gin.Context) {
	var body *types.CreateURLRequest
	err := ctx.BindJSON(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing URL",
		})
		return
	}

	containHttp := strings.Contains(body.Url, "http://")
	containHttps := strings.Contains(body.Url, "https://")

	if !containHttp && !containHttps {
		body.Url = "http://" + body.Url
	}

	BASE_URL := os.Getenv("SHORTEN_BASE_URL")
	id := utils.GenerateNewId(5, body.Url)

	db := utils.GetDB()
	userId, found := ctx.Get("user")

	var uid uuid.UUID
	var uuid_err error

	if found {
		uid, uuid_err = uuid.Parse(userId.(string))
		if uuid_err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid UUID",
			})
			return
		}
	} else {
		uid = uuid.Nil
	}

	create_err := db.Create(&dto.URLShortener{
		Original: body.Url,
		Expires:  body.Expires,
		UrlID:    id,
		UserID:   uid,
	}).Error

	if create_err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to shorten url",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"url": fmt.Sprintf("%v/s/%v", BASE_URL, id),
	})
}

func RedirectToUrl(ctx *gin.Context) {
	id := ctx.Param("id")
	url, ok := utils.GetUrlFromId(id)
	FRONTEND_URL := os.Getenv("FRONTEND_URL")

	if id == "" {
		ctx.Redirect(http.StatusTemporaryRedirect, FRONTEND_URL)
		return
	}

	if !ok {
		ctx.HTML(http.StatusNotFound, "index.html", gin.H{
			"link": FRONTEND_URL,
		})
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func EditOriginalUrl(ctx *gin.Context) {
	var body types.EditURLRequest
	url_short, _ := ctx.Get("post")

	model_url := url_short.(dto.URLShortener)
	bind_err := ctx.BindJSON(&body)

	if bind_err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing New URL",
		})
		return
	}

	containHttp := strings.Contains(body.Url, "http://")
	containHttps := strings.Contains(body.Url, "https://")

	if !containHttp && !containHttps {
		body.Url = "http://" + body.Url
	}

	db := utils.GetDB()

	model_url.Original = body.Url
	model_url.Expires = body.Expires
	db.Save(&model_url)

	ctx.JSON(http.StatusAccepted, gin.H{
		"msg": "Change Successfully",
	})
}

func DeleteShortenerUrl(ctx *gin.Context) {
	url_short, _ := ctx.Get("post")
	model_url := url_short.(dto.URLShortener)

	db := utils.GetDB()
	err := db.Delete(&model_url).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete this id",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete Successfully",
	})
}

func GetHistoryWithLimit(ctx *gin.Context) {
	query_limit := ctx.Query("limit")
	query_page := ctx.Query("pages")

	if query_limit == "" {
		query_limit = "7"
	}

	if query_page == "" {
		query_page = "1"
	}

	limit, err_limit := strconv.Atoi(query_limit)
	pages, err_page := strconv.Atoi(query_page)
	if err_limit != nil || err_page != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Query",
		})
		return
	}

	if pages <= 0 {
		pages = 1
	}

	db := utils.GetDB()
	userId, _ := ctx.Get("user")
	uid, uuid_err := uuid.Parse(userId.(string))
	if uuid_err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID",
		})
		return
	}

	url_list := make([]dto.URLShortener, 0)
	var cnt int64

	db_query := db.Model(&dto.URLShortener{}).Where(dto.URLShortener{UserID: uid})

	tx := db_query.Count(&cnt)
	all_pages := uint(math.Ceil(float64(cnt) / float64(limit)))

	if all_pages <= 1 {
		pages = 1
		all_pages = 1
	}

	if all_pages != 1 && pages >= int(all_pages) {
		pages = int(all_pages)
	}

	err := tx.Where(dto.URLShortener{UserID: uid}).Offset((pages - 1) * limit).Limit(limit).Order("created_at desc").Find(&url_list).Error

	if err != nil || tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something wrong with query",
		})
		return
	}

	res := make([]types.UserHistoryResponse, 0)
	for _, element := range url_list {
		expires, err := utils.CalculateExpiresTime(element.CreatedAt, element.Expires)
		if err == nil {
			res = append(res, types.UserHistoryResponse{
				Original:   element.Original,
				ShortenId:  element.UrlID,
				CreatedAt:  element.CreatedAt,
				Expires:    expires,
				RawExpires: element.Expires,
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":      res,
		"count":     cnt,
		"all_pages": all_pages,
		"current":   pages,
	})
}
