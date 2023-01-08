package controllers

import (
	"net/http"
	"saw/app/services"
	"saw/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	// "github.com/freman/go-steamauth"
)

type LPS struct {
	Login    string
	Password string
	SteamId  string
}

type AccountsController interface {
	GetAccounts(c *gin.Context)
	GetAccount(c *gin.Context)
	AddAccounts(c *gin.Context)
	EditAccounts(c *gin.Context)
	RemoveAccounts(c *gin.Context)
	GetGuardCode(c *gin.Context)
	GetAccountSummaries(c *gin.Context)
	GetLPS(c *gin.Context)
}

func GetAccounts(c *gin.Context) {
	var accounts []models.Account

	models.GetDB().Model(&models.Account{}).
		Find(&accounts, models.Account{Blocked: false})

	if accounts == nil {
		c.AbortWithStatus(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": accounts})
	}
}

func GetAccount(c *gin.Context) {
	var account models.Account

	models.GetDB().Model(&models.Account{}).
		Where("id = ?", c.Param("id")).
		First(&account)

	if account.ID == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": account})
	}
}

func AddAccounts(c *gin.Context) {
	var accounts []models.Account

	if err := c.ShouldBindJSON(&accounts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/// Для повторных данных занимается id, которая по итогу пропускается

	result := models.GetDB().
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "login"}},
			UpdateAll: true,
		}).Create(&accounts)

	c.JSON(http.StatusCreated, gin.H{"data": result.RowsAffected})
}

func UpdateAccount(c *gin.Context) {
	var account models.Account
	var check models.Account

	if err := models.GetDB().First(&check, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	if err := c.ShouldBindJSON(&account); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := models.GetDB().Model(models.Account{}).
		Where("id = ?", c.Param("id")).Updates(&account)

	c.JSON(http.StatusOK, gin.H{"data": result.RowsAffected})
	c.Done()
}

func DeleteAccount(c *gin.Context) {
	var account models.Account

	if err := models.GetDB().First(&account, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	models.GetDB().Delete(&account)
	c.Done()
}

func GetGuardCode(c *gin.Context) {
	var account models.Account

	if err := models.GetDB().First(&account, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	code := services.GetGuardCode(account.SharedSecret)

	if code == "" {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": code})
	}
}

func GetAccountSummaries(c *gin.Context) {
	var account models.Account

	if err := models.GetDB().First(&account, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	code := services.InitAccount(account)

	if code == -1 {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": code})
	}
}

func CheckAvailability(c *gin.Context) {
	var account models.Account

	if err := models.GetDB().First(&account, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	ok := services.Check(account)
	c.JSON(http.StatusOK, gin.H{"data": ok})
}

func GetLPS(c *gin.Context) {
	var lps []LPS

	models.GetDB().Model(&models.Account{}).
		Where("blocked = ?", "false").Find(&lps)

	if lps == nil {
		c.AbortWithStatus(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": lps})
	}
}
