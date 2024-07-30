package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/yelnar0112/project-management/docs"
	"github.com/yelnar0112/project-management/internal/domain"
	"github.com/yelnar0112/project-management/internal/service"
)

// GetProjects godoc
// @Summary Get all projects
// @Description Retrieve a list of all projects
// @Tags projects
// @Produce json
// @Success 200 {array} domain.Entity
// @Failure 500 {object} gin.H{"error": string, "details": string}
// @Router /projects/ [get]
func GetProjects(c *gin.Context) {
	projects, err := service.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project in the system
// @Tags projects
// @Accept json
// @Produce json
// @Param project body domain.Entity true "Project"
// @Success 201 {object} domain.Entity
// @Failure 400 {object} gin.H{"error": string, "details": string}
// @Failure 500 {object} gin.H{"error": string, "details": string}
// @Router /projects/ [post]
func CreateProject(c *gin.Context) {
	var project domain.Entity
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project data", "details": err.Error()})
		return
	}

	project.ID = uuid.New()
	if err := service.CreateProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}

// GetProject godoc
// @Summary Get a project by ID
// @Description Retrieve a project by its ID
// @Tags projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} domain.Entity
// @Failure 400 {object} gin.H{"error": string, "details": string}
// @Failure 404 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string, "details": string}
// @Router /projects/{id} [get]
func GetProject(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID", "details": err.Error()})
		return
	}

	project, err := service.GetProject(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve project", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, project)
}

// UpdateProject godoc
// @Summary Update a project
// @Description Update a project by its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body domain.Entity true "Project"
// @Success 200 {object} domain.Entity
// @Failure 400 {object} gin.H{"error": string, "details": string}
// @Failure 500 {object} gin.H{"error": string, "details": string}
// @Router /projects/{id} [put]
func UpdateProject(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID", "details": err.Error()})
		return
	}

	var project domain.Entity
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project data", "details": err.Error()})
		return
	}

	project.ID = id
	if err := service.UpdateProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

// DeleteProject godoc
// @Summary Delete a project
// @Description Delete a project by its ID
// @Tags projects
// @Param id path string true "Project ID"
// @Success 200 {object} gin.H{"message": string}
// @Failure 400 {object} gin.H{"error": string, "details": string}
// @Failure 500 {object} gin.H{"error": string, "details": string}
// @Router /projects/{id} [delete]
func DeleteProject(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID", "details": err.Error()})
		return
	}

	if err := service.DeleteProject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
