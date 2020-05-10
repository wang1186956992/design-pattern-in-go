package main

func main() {
	req := Request{
		Name:  "课桌椅",
		Price: 60000,
	}

	dh := &DepartmentHandler{
		Handler: Handler{
			Next: nil,
		},
	}

	ch := &CollegeHandler{
		Handler: Handler{
			Next: nil,
		},
	}

	vh := &ViceSchoolHandler{
		Handler: Handler{
			Next: nil,
		},
	}

	sh := &SchoolHandler{
		Handler: Handler{
			Next: nil,
		},
	}

	//构成环状链
	dh.Next = ch
	ch.Next = vh
	vh.Next = sh
	sh.Next = dh

	dh.Handle(req)
}
