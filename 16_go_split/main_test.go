package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var splitTests = []struct {
	name   string
	file   string
	nFile  int
	text   string
	expect []string
}{
	{
		name:  "should split a 5-lines file into 10 pieces",
		file:  "./test.txt",
		nFile: 10,
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920`,
		expect: []string{"1	Caspian_Sea	436,000	78,200",
			"2	Lake_Superior	82,100	12,100",
			"3	Lake_Victoria	68,870	2,760",
			"4	Lake_Huron	59,600	3,540",
			"5	Lake_Michigan	57,800	4,920",
			"",
			"",
			"",
			"",
			""},
	},
	{
		name:  "should split a 24-lines file into 10 pieces",
		file:  "./full-test.txt",
		nFile: 10,
		text: `高知県	江川崎	41	2013-08-12
埼玉県	熊谷	40.9	2007-08-16
岐阜県	多治見	40.9	2007-08-16
山形県	山形	40.8	1933-07-25
山梨県	甲府	40.7	2013-08-10
和歌山県	かつらぎ	40.6	1994-08-08
静岡県	天竜	40.6	1994-08-04
山梨県	勝沼	40.5	2013-08-10
埼玉県	越谷	40.4	2007-08-16
群馬県	館林	40.3	2007-08-16
群馬県	上里見	40.3	1998-07-04
愛知県	愛西	40.3	1994-08-05
千葉県	牛久	40.2	2004-07-20
静岡県	佐久間	40.2	2001-07-24
愛媛県	宇和島	40.2	1927-07-22
山形県	酒田	40.1	1978-08-03
岐阜県	美濃	40	2007-08-16
群馬県	前橋	40	2001-07-24
千葉県	茂原	39.9	2013-08-11
埼玉県	鳩山	39.9	1997-07-05
大阪府	豊中	39.9	1994-08-08
山梨県	大月	39.9	1990-07-19
山形県	鶴岡	39.9	1978-08-03
愛知県	名古屋	39.9	1942-08-02`,
		expect: []string{`高知県	江川崎	41	2013-08-12
埼玉県	熊谷	40.9	2007-08-16
岐阜県	多治見	40.9	2007-08-16`,
			`山形県	山形	40.8	1933-07-25
山梨県	甲府	40.7	2013-08-10
和歌山県	かつらぎ	40.6	1994-08-08`,
			`静岡県	天竜	40.6	1994-08-04
山梨県	勝沼	40.5	2013-08-10
埼玉県	越谷	40.4	2007-08-16`,
			`群馬県	館林	40.3	2007-08-16
群馬県	上里見	40.3	1998-07-04
愛知県	愛西	40.3	1994-08-05`,
			`千葉県	牛久	40.2	2004-07-20
静岡県	佐久間	40.2	2001-07-24`,
			`愛媛県	宇和島	40.2	1927-07-22
山形県	酒田	40.1	1978-08-03`,
			`岐阜県	美濃	40	2007-08-16
群馬県	前橋	40	2001-07-24`,
			`千葉県	茂原	39.9	2013-08-11
埼玉県	鳩山	39.9	1997-07-05`,
			`大阪府	豊中	39.9	1994-08-08
山梨県	大月	39.9	1990-07-19`,
			`山形県	鶴岡	39.9	1978-08-03
愛知県	名古屋	39.9	1942-08-02`},
	},
}

func TestSplit(t *testing.T) {
	for _, testcase := range splitTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}

		f.Write([]byte(testcase.text))
		f.Close()

		if err := split(testcase.file, testcase.nFile); err != nil {
			t.Errorf("could not split a file into pieces: %s\n", err)
		}

		var (
			buf      []byte
			fileName string
		)
		fds := make([]*os.File, testcase.nFile)
		ext := filepath.Ext(testcase.file)
		for i := 0; i < testcase.nFile; i++ {
			fileName = strings.TrimSuffix(testcase.file, ext) + "_" + strconv.Itoa(i+1) + ext
			fds[i], err = os.Open(fileName)
			if err != nil {
				t.Errorf("could not open a file: %s\n  %s\n", fileName, err)
			}
			if buf, _ = ioutil.ReadAll(fds[i]); !reflect.DeepEqual(string(buf), testcase.expect[i]) {
				t.Errorf("result => %#v\n expect => %#v\n", string(buf), testcase.expect[i])
			}
			fds[i].Close()

			if err := os.Remove(fileName); err != nil {
				t.Errorf("could not delete a file: %s\n  %s\n", fileName, err)
			}
		}
		f.Close()

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
