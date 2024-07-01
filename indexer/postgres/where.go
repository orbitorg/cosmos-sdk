package postgres

import (
	"fmt"
	"io"
)

func (tm *TableManager) WhereSqlAndParams(w io.Writer, key interface{}, startParamIdx int) (endParamIdx int, keyParams []interface{}, err error) {
	var keyCols []string
	keyParams, keyCols, err = tm.bindKeyParams(key)
	if err != nil {
		return
	}

	endParamIdx, keyParams, err = tm.WhereSql(w, keyParams, keyCols, startParamIdx)
	return
}

func (tm *TableManager) WhereSql(w io.Writer, params []interface{}, cols []string, startParamIdx int) (endParamIdx int, resParams []interface{}, err error) {
	_, err = fmt.Fprintf(w, " WHERE ")
	if err != nil {
		return
	}

	endParamIdx = startParamIdx
	for i, col := range cols {
		if i > 0 {
			_, err = fmt.Fprintf(w, " AND ")
			if err != nil {
				return
			}
		}

		_, err = fmt.Fprintf(w, "%s ", col)
		if err != nil {
			return
		}

		if params[i] == nil {
			_, err = fmt.Fprintf(w, "IS NULL")
			if err != nil {
				return
			}

		} else {
			_, err = fmt.Fprintf(w, "= $%d", endParamIdx)
			if err != nil {
				return
			}

			resParams = append(resParams, params[i])

			endParamIdx++
		}
	}

	return
}
