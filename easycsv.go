package easycsv

func readCSV(bytes []byte) (map[int]map[string]string, error) {
	csv := make(map[int]map[string]string)
	headers := make(map[int]string)
	var err error

	s := string(bytes)

	off := false
	quote := -1
	numheader := 0
	line := 0
	csv[0] = make(map[string]string)
	for _, i := range s {
		if off == false {
			if i == '\n' && quote != 1 {
				off = true
				numheader = 0
				quote = -1
			} else if i == '"' && quote != 1 {
				quote = 0
			} else if i == '"' && quote == 0 {
				headers[numheader] += "\\" + string(i)
				quote = 1
			} else if i == ',' && quote != 1 {
				numheader++
				if quote == 0 {
					quote = -1
				}
			} else {
				headers[numheader] += string(i)
				if quote == 0 {
					quote = -1
				}
			}
		} else {
			if i == '\n' && quote != 1 {
				numheader = 0
				line++
				csv[line] = make(map[string]string)
				quote = -1
			} else if i == '\n' && quote == 1 {
				csv[line][headers[numheader]] += "\n"
			} else if i == '"' && quote == -1 {
				quote = 2
			} else if i == '"' && quote == 2 {
				quote = -1
			} else if i == '"' && quote != 0 {
				quote = 0
			} else if i == '"' && quote == 0 {
				csv[line][headers[numheader]] += string(i)
				quote = 1
			} else if i == ',' && quote != 1 && quote != 2 {
				numheader++
				if quote == 0 {
					quote = -1
				}
				if quote == 2 {
					quote = 1
				}
			} else {
				csv[line][headers[numheader]] += string(i)
				if quote == 0 {
					quote = -1
				}
				if quote == 2 {
					quote = 1
				}
			}
		}
	}

	return csv, err
}
