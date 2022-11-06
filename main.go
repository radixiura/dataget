package main

import (
  "fmt"
  "os"
)

var p = fmt.Print
var pln = fmt.Println
var fs = fmt.Fscan

func main() {
 pln("Welcome to dataget")
 var surname string
 var name string
 var patronymic string
 var birthdate string
 var birthplace string
 var email string
 var phone string
 var card string
 var papers string
 var paperscode string
 var papersdate string
 var snils string
 var workplace string
 var post string
 var workaddress string
 var workphone string
 
 pln("Фамилия: ")
 pln("Имя: ")
 pln("Отчество: ")
 pln("Дата рождения: ")
 pln("Телефон: ")
 pln("Почта: ")
 pln("Серия и номер паспорта: ")
 pln("Код отделения: ")
 pln("Дата выдачи: ")
 pln("Место рождения: ")
 pln("Снилс: ")
 pln("Место работы: ")
 pln("Должность: ")
 pln("Адрес работы: ")
 pln("Телефон работодателя: ")

 pln("Номер карты: ")
 pln("Дата на карте: ")
 pln("CVV: ")
}
