package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var text1 = `Жил-был у бабушки серенький козлик,
Жил-был у бабушки серенький козлик –
Вот как, вот как серенький козлик,
Вот как, вот как серенький козлик.
Бабушка козлика очень любила,
Бабушка козлика очень любила –
Вот как, вот как очень любила,
Вот как, вот как очень любила!
Вздумалось козлику в лес погуляти,
Вздумалось козлику в лес погуляти –
Вот как, вот как в лес погуляти,
Вот как, вот как в лес погуляти.
Напали на козлика серые волки,
Напали на козлика серые волки –
Вот как, вот как серые волки,
Вот как, вот как серые волки.
Но козлик умчался вниз по дорожке,
Но козлик умчался вниз по дорожке –
Вот как, вот как вниз по дорожке,
Вот как, вот как вниз по дорожке!
Вдали замелькали рожки да ножки,
Вдали замелькали рожки да ножки –
Вот как, вот как рожки да ножки,
Вот как, вот как рожки да ножки.`

var piText = `3 , 1 4 1 5 9 2 6 5 3 5 8 9 7 9 3 2 3 8 4 6 2 6 4 3 3 8 3 2 7 9 5 0 2 8 8 4 1 9 7 1 6 9
	3 9 9 3 7 5 1 0 5 8 2 0 9 7 4 9 4 4 5 9 2 3 0 7 8 1 6 4 0 6 2 8 6 2 0 8 9 9 8 6 2 8 0 3 4 8 2 5 
	3 4 2 1 1 7 0 6 7 9 8 2 1 4 8 0 8 6 5 1 3 2 8 2 3 0 6 6 4 7 0 9 3 8 4 4 6 0 9 5 5 0 5 8 2 2 3 1 
	7 2 5 3 5 9 4 0 8 1 2 8 4 8 1 1 1 7 4 5 0 2 8 4 1 0 2 7 0 1 9 3 8 5 2 1 1 0 5 5 5 9 6 4 4 6 2 2 
	9 4 8 9 5 4 9 3 0 3 8 1 9 6 4 4 2 8 8 1 0 9 7 5 6 6 5 9 3 3 4 4 6 1 2 8 4 7 5 6 4 8 2 3 3 7 8 6 
	7 8 3 1 6 5 2 7 1 2 0 1 9 0 9 1 4 5 6 4 8 5 6 6 9 2 3 4 6 0 3 4 8 6 1 0 4 5 4 3 2 6 6 4 8 2 1 3 
	3 9 3 6 0 7 2 6 0 2 4 9 1 4 1 2 7 3 7 2 4 5 8 7 0 0 6 6 0 6 3 1 5 5 8 8 1 7 4 8 8 1 5 2 0 9 2 0 
	9 6 2 8 2 9 2 5 4 0 9 1 7 1 5 3 6 4 3 6 7 8 9 2 5 9 0 3 6 0 0 1 1 3 3 0 5 3 0 5 4 8 8 2 0 4 6 6 
	5 2 1 3 8 4 1 4 6 9 5 1 9 4 1 5 1 1 6 0 9 4 3 3 0 5 7 2 7 0 3 6 5 7 5 9 5 9 1 9 5 3 0 9 2 1 8 6 
	1 1 7 3 8 1 9 3 2 6 1 1 7 9 3 1 0 5 1 1 8 5 4 8 0 7 4 4 6 2 3 7 9 9 6 2 7 4 9 5 6 7 3 5 1 8 8 5 
	7 5 2 7 2 4 8 9 1 2 2 7 9 3 8 1 8 3 0 1 1 9 4 9 1 2 9 8 3 3 6 7 3 3 6 2 4 4 0 6 5 6 6 4 3 0 8 6 
	0 2 1 3 9 4 9 4 6 3 9 5 2 2 4 7 3 7 1 9 0 7 0 2 1 7 9 8 6 0 9 4 3 7 0 2 7 7 0 5 3 9 2 1 7 1 7 6 
	2 9 3 1 7 6 7 5 2 3 8 4 6 7 4 8 1 8 4 6 7 6 6 9 4 0 5 1 3 2 0 0 0 5 6 8 1 2 7 1 4 5 2 6 3 5 6 0 
	8 2 7 7 8 5 7 7 1 3 4 2 7 5 7 7 8 9 6 0 9 1 7 3 6 3 7 1 7 8 7 2 1 4 6 8 4 4 0 9 0 1 2 2 4 9 5 3 
	4 3 0 1 4 6 5 4 9 5 8 5 3 7 1 0 5 0 7 9 2 2 7 9 6 8 9 2 5 8 9 2 3 5 4 2 0 1 9 9 5 6 1 1 2 1 2 9 
	0 2 1 9 6 0 8 6 4 0 3 4 4 1 8 1 5 9 8 1 3 6 2 9 7 7 4 7 7 1 3 0 9 9 6 0 5 1 8 7 0 7 2 1 1 3 4 9 9 9 9 9 9 `

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})
	t.Run("positive test", func(t *testing.T) {
		expected := []string{
			"он",        // 8
			"а",         // 6
			"и",         // 6
			"ты",        // 5
			"что",       // 5
			"-",         // 4
			"Кристофер", // 4
			"если",      // 4
			"не",        // 4
			"то",        // 4
		}
		require.Equal(t, expected, Top10(text))
	})
	t.Run("simple test for visual control", func(t *testing.T) {
		expected := []string{
			"vr", // ( 5 )
			"aa", // ( 3 )
			"bc", // ( 3 )
		}
		require.Equal(t, expected, Top10("bc bc  bc aa aa   aa vr vr vr vr vr"), 0)
	})
	t.Run("sensei test", func(t *testing.T) {
		expected := []string{
			"a", "c", "d", "j", "k", "m", "n", "p", "q", "r",
		}
		require.Equal(t, expected, Top10("a a v v k k c c d d x x z z n n m m j j p p q q r r "), 0)
	})
	t.Run("another positive test", func(t *testing.T) {
		expected := []string{
			"Вот",     // ( 12 )
			"вот",     // ( 12 )
			"как",     // ( 12 )
			"как,",    // ( 12 )
			"–",       // ( 6 )
			"в",       // ( 4 )
			"вниз",    // ( 4 )
			"да",      // ( 4 )
			"козлика", // ( 4 )
			"лес",     // ( 4 )
		}
		require.Equal(t, expected, Top10(text1))
	})
	t.Run("Chinese test", func(t *testing.T) {
		require.NotPanics(t, func() { Top10("肮 脏 的 灰 色 狐 狸 返 回 宿 舍 一 步 一 步") })
	})
	t.Run("Pi test", func(t *testing.T) {
		expected := []string{
			"1", // ( 88 )
			"9", // ( 85 )
			"2", // ( 81 )
			"4", // ( 79 )
			"3", // ( 76 )
			"6", // ( 75 )
			"8", // ( 72 )
			"0", // ( 71 )
			"7", // ( 71 )
			"5", // ( 70 )
		}
		require.Equal(t, expected, Top10(piText))
	})
}
