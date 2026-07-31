package handlers

import (
	"net/http"
	"stats/entities/dto"
	"stats/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	StatsService services.StatsServiceInterface
}

func NewStatsHandler(statsService services.StatsServiceInterface) *StatsHandler {
	return &StatsHandler{
		StatsService: statsService,
	}
}

func (h *StatsHandler) GetLevelStats(c *gin.Context) {
	ctx := c.Request.Context()

	levelID, err := strconv.Atoi(c.Param("level_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid level_id format",
		})
		return
	}

	levelStats, err := h.StatsService.GetLevelStats(ctx, levelID)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, levelStats)
}

func (h *StatsHandler) CreateLevelStatsEntry(c *gin.Context) {
	ctx := c.Request.Context()

	levelID, err := strconv.Atoi(c.Param("level_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid level_id format",
		})
		return
	}

	var body dto.CreateLevelStatsRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	params := body.ToParams(levelID)

	statsEntry, err := h.StatsService.CreateLevelStatsEntry(ctx, params)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, statsEntry)
}

func (h *StatsHandler) GetPlayerRecords(c *gin.Context) {
	ctx := c.Request.Context()

	username := c.Param("username")

	playerRecords, err := h.StatsService.GetPlayerRecords(ctx, username)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, playerRecords)
}
