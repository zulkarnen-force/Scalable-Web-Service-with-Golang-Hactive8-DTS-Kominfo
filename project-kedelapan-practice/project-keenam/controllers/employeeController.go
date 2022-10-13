package controllers

import (
	"net/http"
	"strconv"

	"project-keenam/models"
	// "project-keenam/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

// CreateEmployee godoc
// @Summary     create an employee
// @Description create an employee
// @Tags        employees
// @Accept      json
// @Produce     json
// @Param		request body models.Employee true "create employee"
// @Success     201 {object} models.Employee
// @Router      /employees [post]
func (idb *InDB) CreateEmployee (ctx *gin.Context) {
		var newEmployee models.Employee
		ctx.ShouldBindJSON(&newEmployee)

		idb.DB.Debug().Create(&newEmployee)

		ctx.JSON(http.StatusCreated, newEmployee)
}


// GetAllEmployee godoc
// @Summary     list all of employee
// @Description list all of employee
// @Tags        employees
// @Accept      json
// @Produce     json
// @Success     200 {object} []models.Employee
// @Router      /employees [get]
func (idb *InDB) GetAllEmployees (ctx *gin.Context) {
		var employees []models.Employee

		idb.DB.Debug().Find(&employees)

		ctx.JSON(http.StatusOK, employees)
	}

func (idb *InDB) GetAllEmployeesWithView (ctx *gin.Context)  {
		var employees []models.Employee

		idb.DB.Debug().Find(&employees)
		ctx.HTML(http.StatusOK, "template.html", employees)
}

// ShowEmployee godoc
// @Summary     show employee
// @Description show employee
// @Tags        employees
// @Accept      json
// @Produce     json
// @Param 		employee_id path string true "employee_id"
// @Success     200 {object} models.Employee
// @Router      /employees/{employee_id} [get]
func (idb *InDB) ShowEmployee (ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		idEmployee, _ := strconv.Atoi(id)

		var employee models.Employee

		idb.DB.Debug().First(&employee, idEmployee)

		if employee.ID == 0 {
			ctx.JSON(http.StatusNotFound, "Data Not Found")
			return
		}

		ctx.JSON(http.StatusOK, employee)
}


// DeleteEmployee godoc
// @Summary     delete employee
// @Description delete employee
// @Tags        employees
// @Accept      json
// @Produce     json
// @Param 		employee_id path string true "employee_id"
// @Success     200 {object} models.Employee
// @Router      /employees/{employee_id} [delete]
func (idb *InDB) DeleteEmployee(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		idEmployee, _ := strconv.Atoi(id)

		var employee models.Employee
		idb.DB.Debug().Delete(&employee, idEmployee)

		ctx.JSON(http.StatusOK, "Successfuly Deleting Data")
}


// UpdateEmployee godoc
// @Summary     update employee
// @Description update employee
// @Tags        employees
// @Accept      json
// @Produce     json
// @Param 		employee_id path string true "employee_id"
// @Param 		request body models.Employee true "request"
// @Success     200 {object} models.Employee
// @Router      /employees/{employee_id} [put]
func (idb *InDB) UpdateEmployee (ctx *gin.Context) {
		var updateEmployee models.Employee
		ctx.ShouldBindJSON(&updateEmployee)

		id, _ := ctx.Params.Get("id")
		idEmployee, _ := strconv.Atoi(id)

		idb.DB.Debug().Model(&updateEmployee).Where("id = ?", idEmployee).Updates(updateEmployee)


		ctx.JSON(http.StatusOK, "Successfuly Updating Data")
}