package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
// var taskWithAsteriskIsCompleted = false

// var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
// 	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
// 	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
// 	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
// 		кажется, что можно бы найти какой-то другой способ, если бы  он
// 	только   мог   на  минутку  перестать  бумкать  и  как  следует
// 	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
// 		Как бы то ни было, вот он уже спустился  и  готов  с  вами
// 	познакомиться.
// 	- Винни-Пух. Очень приятно!
// 		Вас,  вероятно,  удивляет, почему его так странно зовут, а
// 	если вы знаете английский, то вы удивитесь еще больше.
// 		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
// 	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
// 	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
// 	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
// 	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
// 	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
// 	звал его тихо, то все подумают, что ты  просто  подул  себе  на
// 	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
// 	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
// 	зря.
// 		А  Винни - так звали самую лучшую, самую добрую медведицу
// 	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
// 	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
// 	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
// 	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
// 	забыл.
// 		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
// 		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
// 	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
// 	посидеть у огня и послушать какую-нибудь интересную сказку.
// 		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	// tests := []struct {
	// 	name     string
	// 	text     string
	// 	expected []string
	// }{
	// 	{
	// 		name:     "Empty input",
	// 		text:     "",
	// 		expected: []string{},
	// 	},
	// 	{
	// 		name:     "Single word",
	// 		text:     "word",
	// 		expected: []string{"word"},
	// 	},
	// 	{
	// 		name:     "All the same word",
	// 		text:     "repeat repeat repeat repeat repeat",
	// 		expected: []string{"repeat"},
	// 	},
	// 	{
	// 		name:     "Each word is unique",
	// 		text:     "each word is unique",
	// 		expected: []string{"each", "is", "unique", "word"},
	// 	},
	// 	{
	// 		name:     "Punctuation and case sensitivity",
	// 		text:     "Hello, hello, HELLO, he'll, he'll, he'll",
	// 		expected: []string{"he'll,", "HELLO,", "Hello,", "he'll", "hello,"},
	// 	},
	// 	{
	// 		name:     "Words with same frequency",
	// 		text:     "one one two two three three four four five five",
	// 		expected: []string{"five", "four", "one", "three", "two"},
	// 	},
	// 	{
	// 		name:     "More than ten unique words",
	// 		text:     "this test includes more than ten unique words each word appears only once",
	// 		expected: []string{"appears", "each", "includes", "more", "once", "only", "ten", "test", "than", "this"},
	// 	},
	// 	{
	// 		name:     "Words with hyphens and apostrophes",
	// 		text:     "it's the time-tested method-of-operation for someone's specific-case",
	// 		expected: []string{"for", "it's", "method-of-operation", "someone's", "specific-case", "the", "time-tested"},
	// 	},
	// }
	// for _, test := range tests {
	// 	t.Run(test.name, func(t *testing.T) {
	// 		require.Equal(t, test.expected, Top10(test.text))
	// 	})
	// }

	// t.Run("positive test", func(t *testing.T) {
	// 	if taskWithAsteriskIsCompleted {
	// 		expected := []string{
	// 			"а",         // 8
	// 			"он",        // 8
	// 			"и",         // 6
	// 			"ты",        // 5
	// 			"что",       // 5
	// 			"в",         // 4
	// 			"его",       // 4
	// 			"если",      // 4
	// 			"кристофер", // 4
	// 			"не",        // 4
	// 		}
	// 		require.Equal(t, expected, Top10(text))
	// 	} else {
	// 		expected := []string{
	// 			"он",        // 8
	// 			"а",         // 6
	// 			"и",         // 6
	// 			"ты",        // 5
	// 			"что",       // 5
	// 			"-",         // 4
	// 			"Кристофер", // 4
	// 			"если",      // 4
	// 			"не",        // 4
	// 			"то",        // 4
	// 		}
	// 		require.Equal(t, expected, Top10(text))
	// 	}
	// })

	testsExtra := []struct {
		name     string
		text     string
		expected []string
	}{
		{
			name: "Mixed case with punctuation",
			text: `This morning, I asked my colleague: 'Have you seen the reports 
			from yesterday?' They replied: 'Yes, I saw the Reports yesterday evening.' 
			The dialogue went on while we discussed various topics, including the upcoming 
			meeting and the latest news from the office. When I returned to my desk, I began to draft the meeting notes, 
			making sure everything was recorded properly. By the end of the day, 
			I was exhausted from all the talking and writing.`,
			expected: []string{"the", "i", "from", "and", "meeting", "my", "reports", "to", "was", "yesterday"},
		},
		{
			name: "Text with apostrophes and hyphens",
			text: `It's important to remember that well-being isn't just about physical health; it's about mental health too. 
			Today's seminar focuses on how we're adapting to new challenges in today's fast-paced world. 
			We'll discuss stress-management techniques and how they're applied in our daily lives. 
			Remember, it's not just about handling stress—it's about thriving in spite of it.`,
			expected: []string{"about", "in", "it's", "health", "how", "just", "remember", "to", "today's", "adapting"},
		},
		{
			name: "Repeated words with different punctuation and case",
			text: `Hello, hello, HELLO! How are you doing today? Hello again, my friend. 
			Today is a wonderful day to say hello to everyone you meet. Hello, hello, HELLO!`,
			expected: []string{"hello", "to", "today", "you", "a", "again", "are", "day", "doing", "everyone"},
		},
	}

	for _, test := range testsExtra {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, Top10(test.text))
		})
	}
}
