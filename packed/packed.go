package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBgYGDoMv0YwYAElBk4GRJTUvLzivXTMnNSQwty8hNT9ItSi/NLi5JT9TPzMkuyivPzQkNYGRhdMqanBXizcyDrh5nMgWGyGXEm6yclFqfGF1cWx+em5pXqoVh25drBvikGEuwfftu4Xv5s4czNs9S8v9pQbZ16so/x37TtG7ewnMrp+fgs09+uya6ls7B2zsUfEm0MfoEiTkkxfpN//thXz2mf8Sun9cJrUSeORFVD9xkbFFJ7Ls7YoKDmeaZJZP3stlrGQweXnlqQ57Pt5DQ2EyeVh4ZPH57fpjW9Povbktd8d0vZWiETiwn6LYdlZnq93javm6ln94S/qwR9n3zvz99fW1f+q6T616d93lo5Txf8UGXtWvNm+p3Ve/o4bu3PyvedG/HfzqJjse4um2gPDblQyWcdW5Yubpp9/FDsd+F4pif8W62arq88pX7927H45XsMlgpHVW/l8Sm9eYdPLCiqOZxjbyR/1RKT1EAptsAZk2SlPyfMEfC7qdZ1csZJh9csk/fvVZENOly8a/fvqt8PzVX22Grw+hfntt55t9ZyhbZNl1vbhW9mp9fMXNZmrK3Nu3R/Re3OnPzE4jUM/omqzGtSzLS9Pee4mcw6s6tbffO8aPXNfPuzZVdPKDql3n6+f4Y5//Tqyboab5urDOclG/x//Vvd56PSjFkzZgb9faia/FascemZg+eWfHKfzjNVh8lr9aWzLSC26d8PH+0XFhWaln/+DOJWZe0XDzjbZbxy1sOvTuVHr36WZ2Rg+P8flIJsXdSMrzIyMNTwwlIQCExGS0HS+FMQOJmEgdMkLHWjmyCF1wQkA7AnakwDBbEZSIY5bHBz8GhmZBJhxp1rIUCA4a0jiCYtD8NMhuRa5Ngwg5vMwLCkMYo4k/HlYZhluCIY5o3/jnOYGIiMbkTAYItwhIn38JpIlIGCKAbqMzPgjX7c5rChmFMONwdJMysbRCEbw3tGBoYtzCAeIAAA///KdhXhuQUAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
