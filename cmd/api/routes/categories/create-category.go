package categories

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	c_repository "github.com/GabrielBrotas/go-categories-msvc/internal/categories/repository"
	c_use_cases "github.com/GabrielBrotas/go-categories-msvc/internal/categories/use-cases"
)

type createCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Teste       string `json:"teste"`
}

// @Summary  Criar categoria
// @Tags     categories
// @Accept   json
// @Produce  json
// @Param    body  body      createCategoryInput  true  "Nome da categoria"
// @Success  201   {object}  map[string]interface{}
// @Failure  400   {object}  map[string]interface{}
// @Router   /categories [post]
func createCategory(c *gin.Context, repository *c_repository.CategoryRepository) {

	var requestBody createCategoryInput

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	if len(requestBody.Name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid Name"})
		return
	}

	use_case := c_use_cases.NewCreateCategoryUseCase(repository)
	err := use_case.Execute(requestBody.Name, requestBody.Description, requestBody.Teste)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true})
}
