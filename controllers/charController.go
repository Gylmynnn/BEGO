package controllers

import (
	"errors"
	"net/http"
	"trygo/rest-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetData(c *gin.Context) {
  var datas []models.Character
  models.DB.Find(&datas)


  c.JSON(200, models.ResType{
    Success: true,
    Message: "successfully",
    Data: datas,
  })
}


func GetErrorMsg(fe validator.FieldError) string {
   switch fe.Tag() {
  case "required" :
        return "field required"
  }
  return "Unknown Error"
}

func GetDataById (c *gin.Context) {
  var datas models.Character
  if err := models.DB.Where("id = ?", c.Param("id")).First(&datas).Error; err != nil {
    c.JSON(http.StatusBadRequest, models.ResType{
      Success: false,
      Message: "data not found",
      Data: err,
    })
  }

  c.JSON(200, models.ResType{
    Success: true,
    Message: "successfully",
    Data: datas,
  })
}


func PostData(c *gin.Context) {
  var input models.ValidatePostInput
  if err := c.ShouldBindJSON(&input); err != nil {
    var ve validator.ValidationErrors
    if errors.As(err, &ve) {
      out := make([]models.ErrorMsg, len(ve))
      for i, fe := range ve {
        out[i] = models.ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
      }
      c.AbortWithStatusJSON(http.StatusBadRequest, models.ResType{
        Success: false,
        Message: "error",
        Data: out,
      })
    }
    return
  }

  datas := models.Character{
    Character: input.Character,
    Quote: input.Quote,
  }

  models.DB.Create(&datas)

  c.JSON(201, models.ResType{
    Success: true,
    Message: "successfully",
    Data: datas,
  })
}


func PutData(c*gin.Context) {

  var datas models.Character
  if err := models.DB.Where("id = ?", c.Param("id")).First(&datas).Error; err != nil {
    c.JSON(http.StatusBadRequest, models.ResType{
      Success: false,
      Message: "data not found",
      Data: err,
    })
  }


  var input models.ValidatePostInput
  if err := c.ShouldBindJSON(&input); err != nil {
    var ve validator.ValidationErrors
    if errors.As(err, &ve) {
      out := make([]models.ErrorMsg, len(ve))
      for i, fe := range ve {
        out[i] = models.ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
      }
      c.AbortWithStatusJSON(http.StatusBadRequest, models.ResType{
        Success: false,
        Message: "error",
        Data: out,
      })
    }
    return
  }

  models.DB.Model(&datas).Updates(input)

  c.JSON(200, models.ResType{
    Success: true,
    Message: " update successfully",
    Data: datas,
  })

}


func DeleteData(c*gin.Context) {
    var datas models.Character
  if err := models.DB.Where("id = ?",c.Query("id")).First(&datas).Error; err != nil {
    c.JSON(http.StatusBadRequest, models.ResType{
      Success: false,
      Message: "data not found",
      Data: err,
    })
  }

  models.DB.Delete(&datas)

  c.JSON(200, models.ResType{
    Success: true,
    Message: " delete successfully",
    Data: datas,
  })


}

