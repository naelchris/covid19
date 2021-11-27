package class

// func AddCasesHandler(w http.ResponseWriter, r *http.Request) {
// 	timeStart := time.Now()
// 	var (
// 		ctx  = r.Context()
// 		re covid.Cases
// 	)

// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		log.Println("[classHandler][AddClassHandler][jsonNewDocoder] decode error err, ", err)
// 		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
// 		return
// 	}

// 	res, err := server.ClassUsecase.AddClass(ctx, data)
// 	if err != nil {
// 		log.Println("[classHandler][AddClassHandler] unable to read body err :", err)
// 		server.RenderError(w, http.StatusBadRequest, err, timeStart)
// 		return
// 	}

// 	resp := classResponse{
// 		ID: res.ClassID,
// 	}

// 	server.RenderResponse(w, http.StatusCreated, resp, timeStart)
// }
