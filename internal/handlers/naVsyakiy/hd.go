package naVsyakiy

// Возвращает данные клиенту
//(func GetHandler(c echo.Context) error {
//	var task []taskService.Task
//	if err := DB.Find(&task).Error; err != nil {
//		return c.JSON(http.StatusBadRequest, taskService.Response{
//			Status:  "error",
//			Message: "Could not add the message",
//		})
//	}
//return c.JSON(http.StatusOK, task)
//}

// Принимает данные от клиента
//func PostHandler(c echo.Context) error {
//	var task taskService.Task
//	if err := c.Bind(&task); err != nil {
//		return c.JSON(http.StatusBadRequest, taskService.Response{
//			Status:  "Error",
//			Message: "Could not add the message",
//		})
//	}
//	if err := DB.Create(&task).Error; err != nil {
//		return c.JSON(http.StatusBadRequest, taskService.Response{
//			Status:  "Error",
//			Message: "Could not add the message",
//		})
//	}
//	return c.JSON(http.StatusOK, task)
//}

// Обновляет данные по ID
//func PatchHandler(c echo.Context) error {
//	idParam := c.Param("id")
//	id, err := strconv.Atoi(idParam)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, taskService.Response{
//		Status:  "Error",
//		Message: "Bad ID",
//		})
//	}

//	var task taskService.Task
//	var updatedMessage taskService.Task
//	if err := c.Bind(&updatedMessage); err != nil {
//		return c.JSON(http.StatusBadRequest, taskService.Response{
//			Status:  "Error",
//			Message: "Could not update the message",
//		})
//	}
// Обновляем задачу в базе данных
//	if err := DB.Model(&taskService.Task{}).Where("id = ?", id).Update("task", updatedMessage.Message).Error; err != nil {
//		return c.JSON(http.StatusBadRequest, taskService.Response{
//			Status:  "Error",
//			Message: "Could not update the message",
//		})
//	}

// Получаем обновленную задачу из базы данных
//	if err := DB.First(&task, id).Error; err != nil {
//	return c.JSON(http.StatusBadRequest, taskService.Response{
//			Status:  "Error",
//			Message: "Could not find the updated task",
//		})
//	}
// Возвращаем обновленную задачу
//	return c.JSON(http.StatusOK, task)
//}

// Удаляет данные по ID
//func DeleteHandler(c echo.Context) error {
//	idParam := c.Param("id")
//	id, err := strconv.Atoi(idParam)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, taskService.Response{
//		Status:  "Error",
//			Message: "Bad ID",
//		})
//	}

//	if err := DB.Delete(&taskService.Task{}, id).Error; err != nil {
//		return c.JSON(http.StatusBadRequest, taskService.Response{
//			Status:  "Error",
//			Message: "Could not delete the message",
//	})
//	}
//	return c.NoContent(http.StatusNoContent)
//}
//
