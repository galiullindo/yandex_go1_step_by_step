package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func getDate(reader *bufio.Reader, layout string) (time.Time, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return time.Time{}, err
	}

	stringDate := strings.TrimSpace(line)
	date, err := time.Parse(layout, stringDate)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func getText(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	return strings.TrimSpace(line), nil
}

func getFloat(reader *bufio.Reader) (float64, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return 0, err
	}

	number, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func getMessage(fam string, im string, otch string, date time.Time, rub int, kop int) string {
	/*
	   Уважаемый, Иванов Андрей Валерьевич, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.
	   Дата подписания договора: 25.04.2005. Просим вас подойти в офис в любое удобное для вас время в этот день.
	   Общая сумма выплат составит 35122 руб. 0 коп.

	   С уважением,
	   Гл. бух. Иванов А.Е.
	*/

	var message string
	message += fmt.Sprintf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n", fam, im, otch)
	message += fmt.Sprintf("Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\n", date.Format("02.01.2006"))
	message += fmt.Sprintf("Общая сумма выплат составит %d руб. %d коп.\n\n", rub, kop)
	message += "С уважением,\nГл. бух. Иванов А.Е.\n"

	return message
}

func check(err error) {
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
	}
}

func main() {
	layout := "02.01.2006"
	reader := bufio.NewReader(os.Stdin)

	date, err := getDate(reader, layout)
	check(err)
	msgDate := date.AddDate(0, 0, 15)

	im, err := getText(reader)
	check(err)
	fam, err := getText(reader)
	check(err)
	otch, err := getText(reader)
	check(err)

	var summa float64
	for range 3 {
		number, err := getFloat(reader)
		check(err)

		summa += number
	}

	rub := int(summa)
	kop := int((summa - float64(rub)) * 100)

	message := getMessage(fam, im, otch, msgDate, rub, kop)
	fmt.Println(message)

}
