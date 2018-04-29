package main

import (
	"bufio"
	"os"
	"testing"
)

var uniqTests = []struct {
	name   string
	file   string
	nRow   int
	text   string
	dest   string
	expect []string
}{
	{
		name: "should filter repeated lines of first colum in a file",
		file: "./test.txt",
		nRow: 1,
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
		dest: "./result.txt",
		expect: []string{"高知県",
			"埼玉県",
			"岐阜県",
			"山形県",
			"山梨県",
			"和歌山県",
			"静岡県",
			"群馬県",
			"愛知県",
			"千葉県",
			"愛媛県",
			"大阪府"},
	},
	{
		name: "should print lines of second colum in a file without any need to filter",
		file: "./test.txt",
		nRow: 2,
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920`,
		dest: "./result.txt",
		expect: []string{"Caspian_Sea",
			"Lake_Superior",
			"Lake_Victoria",
			"Lake_Huron",
			"Lake_Michigan"},
	},
	{
		name: "should return empty string while input nRow is over the length of rows on a file",
		file: "./test.txt",
		nRow: 10,
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920`,
		dest:   "./result.txt",
		expect: []string{},
	},
}

func TestUniq(t *testing.T) {
	for _, testcase := range uniqTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}

		f.Write([]byte(testcase.text))
		f.Close()

		dest, _ := os.Create(testcase.dest)
		if err = uniq(testcase.file, testcase.nRow, *dest); err != nil {
			t.Errorf("could not filter a line in specified file: %s\n  %s\n", testcase.file, err)
		}
		dest.Close()

		dest, _ = os.Open(testcase.dest)
		sc := bufio.NewScanner(dest)
		result := Item{}
		for sc.Scan() {
			result.Add(sc.Text())
		}
		dest.Close()

		for _, expect := range testcase.expect {
			if !result.Contains(expect) {
				t.Errorf("result => %#v\n should contain => %#v\n", result, expect)
			}
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
		if err := os.Remove(testcase.dest); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.dest, err)
		}
	}
}
