package helpers

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewIncrementId(db *pgxpool.Pool, tableName string, prifix string, Length int) (func() string, error) {
	var (
		id sql.NullString
		// capitalLetter = string(strings.ToUpper(tableName)[0])
		query = fmt.Sprintf("SELECT increment_id FROM %s ORDER BY created_at DESC LIMIT 1", tableName)
	)
	contex, cancelF := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*2))

	resp := db.QueryRow(contex, query)

	resp.Scan(&id)
	defer cancelF()

	if !id.Valid {
		fmt.Printf("$$$$$$$$$$$$$$     %+v     $$$$$$$$$$$$$$", id)
		id.String = ""
	}
	return func() string {
		idNumber := idToInt(id.String)
		idNumber++
		var (
			numberLenght = idNumber
			count        = 0
		)
		for numberLenght > 0 {
			numberLenght /= 10
			count++
		}
		if count == 0 {
			count++
		}
		fmt.Printf("COUNT %d  Length %d  &&&&&&&&&&&\n", count, idNumber)

		id := fmt.Sprintf("%s-%s%d", prifix, strings.Repeat("O", Length-count), idNumber)
		return id

	}, nil
}

func idToInt(DatabaseId string) int {
	pattern := regexp.MustCompile("[0-9]+")
	firstMatchSubstring := pattern.FindString(DatabaseId)
	id, err := strconv.Atoi(firstMatchSubstring)
	if err != nil {
		return 0
	}
	fmt.Println("#####################@@@@@@  ", id, "   ", firstMatchSubstring, "  @@@@@@####################")

	return id
}
