package models

    type ResType struct {
        Success bool        `json:"success"`
        Message string      `json:"message"`
        Data    interface{} `json:"data"`
    }


      type ValidatePostInput struct {
	      Character   string `json:"character" binding:"required"`
	      Quote string `json:"quote" binding:"required"`
    }

      type ErrorMsg struct {
	      Field   string `json:"field"`
	      Message string `json:"message"`
    }
