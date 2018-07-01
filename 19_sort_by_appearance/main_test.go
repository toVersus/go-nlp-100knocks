package main

import (
	"reflect"
	"strings"
	"testing"
)

var sortByOverlapCountTests = []struct {
	name      string
	columnNum int
	text      string
	expect    string
}{
	{
		name:      "should sort the full text by counting overlaps of 1st column",
		columnNum: 1,
		text: `高知県	江川崎	41	2013-08-12
埼玉県	熊谷	40.9	2007-08-16
岐阜県	多治見	40.9	2007-08-16
山梨県	甲府	40.7	2013-08-10
山形県	山形	40.8	1933-07-25
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
		expect: `群馬県	館林	40.3	2007-08-16
群馬県	上里見	40.3	1998-07-04
群馬県	前橋	40	2001-07-24
山梨県	甲府	40.7	2013-08-10
山梨県	勝沼	40.5	2013-08-10
山梨県	大月	39.9	1990-07-19
山形県	山形	40.8	1933-07-25
山形県	酒田	40.1	1978-08-03
山形県	鶴岡	39.9	1978-08-03
埼玉県	熊谷	40.9	2007-08-16
埼玉県	越谷	40.4	2007-08-16
埼玉県	鳩山	39.9	1997-07-05
静岡県	天竜	40.6	1994-08-04
静岡県	佐久間	40.2	2001-07-24
愛知県	愛西	40.3	1994-08-05
愛知県	名古屋	39.9	1942-08-02
岐阜県	多治見	40.9	2007-08-16
岐阜県	美濃	40	2007-08-16
千葉県	牛久	40.2	2004-07-20
千葉県	茂原	39.9	2013-08-11
高知県	江川崎	41	2013-08-12
愛媛県	宇和島	40.2	1927-07-22
大阪府	豊中	39.9	1994-08-08
和歌山県	かつらぎ	40.6	1994-08-08`,
	},
}

func TestSortByOverlapCount(t *testing.T) {
	for _, testcase := range sortByOverlapCountTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		result := sortByVotes(r, testcase.columnNum).String()
		if !reflect.DeepEqual(string(result), testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", string(result), testcase.expect)
		}
	}
}
