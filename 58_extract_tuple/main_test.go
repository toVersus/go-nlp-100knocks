package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

var collapsedDepTests = []struct {
	name   string
	text   string
	expect string
}{
	{
		name: "should slice the simple sentence into words",
		text: `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet href="CoreNLP-to-HTML.xsl" type="text/xsl"?>
<root>
	<document>
	<docId>test.txt</docId>
	<sentences>
		<sentence id="1">
		<tokens>
			<token id="1">
			<word>Many</word>
			<lemma>many</lemma>
			<CharacterOffsetBegin>0</CharacterOffsetBegin>
			<CharacterOffsetEnd>4</CharacterOffsetEnd>
			<POS>JJ</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="2">
			<word>challenges</word>
			<lemma>challenge</lemma>
			<CharacterOffsetBegin>5</CharacterOffsetBegin>
			<CharacterOffsetEnd>15</CharacterOffsetEnd>
			<POS>NNS</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="3">
			<word>in</word>
			<lemma>in</lemma>
			<CharacterOffsetBegin>16</CharacterOffsetBegin>
			<CharacterOffsetEnd>18</CharacterOffsetEnd>
			<POS>IN</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="4">
			<word>NLP</word>
			<lemma>nlp</lemma>
			<CharacterOffsetBegin>19</CharacterOffsetBegin>
			<CharacterOffsetEnd>22</CharacterOffsetEnd>
			<POS>NN</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="5">
			<word>involve</word>
			<lemma>involve</lemma>
			<CharacterOffsetBegin>23</CharacterOffsetBegin>
			<CharacterOffsetEnd>30</CharacterOffsetEnd>
			<POS>VBP</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="6">
			<word>natural</word>
			<lemma>natural</lemma>
			<CharacterOffsetBegin>31</CharacterOffsetBegin>
			<CharacterOffsetEnd>38</CharacterOffsetEnd>
			<POS>JJ</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="7">
			<word>language</word>
			<lemma>language</lemma>
			<CharacterOffsetBegin>39</CharacterOffsetBegin>
			<CharacterOffsetEnd>47</CharacterOffsetEnd>
			<POS>NN</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="8">
			<word>understanding</word>
			<lemma>understanding</lemma>
			<CharacterOffsetBegin>48</CharacterOffsetBegin>
			<CharacterOffsetEnd>61</CharacterOffsetEnd>
			<POS>NN</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="9">
			<word>,</word>
			<lemma>,</lemma>
			<CharacterOffsetBegin>61</CharacterOffsetBegin>
			<CharacterOffsetEnd>62</CharacterOffsetEnd>
			<POS>,</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="10">
			<word>that</word>
			<lemma>that</lemma>
			<CharacterOffsetBegin>63</CharacterOffsetBegin>
			<CharacterOffsetEnd>67</CharacterOffsetEnd>
			<POS>WDT</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="11">
			<word>is</word>
			<lemma>be</lemma>
			<CharacterOffsetBegin>68</CharacterOffsetBegin>
			<CharacterOffsetEnd>70</CharacterOffsetEnd>
			<POS>VBZ</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="12">
			<word>,</word>
			<lemma>,</lemma>
			<CharacterOffsetBegin>70</CharacterOffsetBegin>
			<CharacterOffsetEnd>71</CharacterOffsetEnd>
			<POS>,</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="13">
			<word>enabling</word>
			<lemma>enable</lemma>
			<CharacterOffsetBegin>72</CharacterOffsetBegin>
			<CharacterOffsetEnd>80</CharacterOffsetEnd>
			<POS>VBG</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="14">
			<word>computers</word>
			<lemma>computer</lemma>
			<CharacterOffsetBegin>81</CharacterOffsetBegin>
			<CharacterOffsetEnd>90</CharacterOffsetEnd>
			<POS>NNS</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="15">
			<word>to</word>
			<lemma>to</lemma>
			<CharacterOffsetBegin>91</CharacterOffsetBegin>
			<CharacterOffsetEnd>93</CharacterOffsetEnd>
			<POS>TO</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="16">
			<word>derive</word>
			<lemma>derive</lemma>
			<CharacterOffsetBegin>94</CharacterOffsetBegin>
			<CharacterOffsetEnd>100</CharacterOffsetEnd>
			<POS>VB</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="17">
			<word>meaning</word>
			<lemma>meaning</lemma>
			<CharacterOffsetBegin>101</CharacterOffsetBegin>
			<CharacterOffsetEnd>108</CharacterOffsetEnd>
			<POS>NN</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="18">
			<word>.</word>
			<lemma>.</lemma>
			<CharacterOffsetBegin>108</CharacterOffsetBegin>
			<CharacterOffsetEnd>109</CharacterOffsetEnd>
			<POS>.</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
		</tokens>
		<parse>(ROOT (S (NP (NP (JJ Many) (NNS challenges)) (PP (IN in) (NP (NN NLP)))) (VP (VBP involve) (S (NP (NP (JJ natural) (NN language) (NN understanding)) (, ,) (SBAR (WHNP (WDT that)) (S (VP (VBZ is)))) (, ,)) (VP (VBG enabling) (NP (NNS computers)) (S (VP (TO to) (VP (VB derive) (NP (NN meaning)))))))) (. .))) </parse>
		<dependencies type="collapsed-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="5">involve</dependent>
			</dep>
			<dep type="amod">
			<governor idx="2">challenges</governor>
			<dependent idx="1">Many</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="5">involve</governor>
			<dependent idx="2">challenges</dependent>
			</dep>
			<dep type="case">
			<governor idx="4">NLP</governor>
			<dependent idx="3">in</dependent>
			</dep>
			<dep type="nmod:in">
			<governor idx="2">challenges</governor>
			<dependent idx="4">NLP</dependent>
			</dep>
			<dep type="amod">
			<governor idx="8">understanding</governor>
			<dependent idx="6">natural</dependent>
			</dep>
			<dep type="compound">
			<governor idx="8">understanding</governor>
			<dependent idx="7">language</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="13">enabling</governor>
			<dependent idx="8">understanding</dependent>
			</dep>
			<dep type="punct">
			<governor idx="8">understanding</governor>
			<dependent idx="9">,</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="11">is</governor>
			<dependent idx="10">that</dependent>
			</dep>
			<dep type="acl:relcl">
			<governor idx="8">understanding</governor>
			<dependent idx="11">is</dependent>
			</dep>
			<dep type="punct">
			<governor idx="8">understanding</governor>
			<dependent idx="12">,</dependent>
			</dep>
			<dep type="dep">
			<governor idx="5">involve</governor>
			<dependent idx="13">enabling</dependent>
			</dep>
			<dep type="dobj">
			<governor idx="13">enabling</governor>
			<dependent idx="14">computers</dependent>
			</dep>
			<dep type="mark">
			<governor idx="16">derive</governor>
			<dependent idx="15">to</dependent>
			</dep>
			<dep type="advcl">
			<governor idx="13">enabling</governor>
			<dependent idx="16">derive</dependent>
			</dep>
			<dep type="dobj">
			<governor idx="16">derive</governor>
			<dependent idx="17">meaning</dependent>
			</dep>
			<dep type="punct">
			<governor idx="5">involve</governor>
			<dependent idx="18">.</dependent>
			</dep>
		</dependencies>
		</sentence>
	</sentences>
	<coreference/>
	</document>
</root>`,
		expect: "understanding	enabling	computers",
	},
}

func TestCollapsedDep(t *testing.T) {
	for _, testcase := range collapsedDepTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		root, err := readXML(r)
		if err != nil {
			t.Error(err)
		}

		var buf bytes.Buffer
		for _, dep := range root.Document.Sentences[0].Dependencies {
			verbs := dep.getVerb()
			if len(verbs) == 0 {
				continue
			}

			subjects := dep.getSubject(verbs)
			objects := dep.getObject(verbs)

			for i := 0; i < len(verbs); i++ {
				buf.WriteString(subjects[i] + "\t" + verbs[i] + "\t" + objects[i] + "\n")
			}
		}
		result := strings.TrimRight(buf.String(), "\n")
		if diff := deep.Equal(result, testcase.expect); diff != nil {
			t.Error(diff)
		}
	}
}
