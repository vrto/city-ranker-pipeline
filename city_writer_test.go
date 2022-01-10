package parallel_mapper

import "testing"

func TestCityWriter_Write_NothingToWrite(t *testing.T) {
	t.Log("Given CityWriter with a mocked appender.")
	{
		appender := mockAppender{}
		_ = NewCityWriter("test", &appender)
		t.Log("When there is nothing to write.")
		{
			// no Write call invoked, just the constructor
			if appender.captured == "=== test ===\n" {
				t.Log("Then writer only creates an empty header.")
			} else {
				t.Errorf("Then writer should've created only an empty header, but got: %s", appender.captured)
			}
		}
	}
}

func TestCityWriter_Write(t *testing.T) {
	t.Log("Given CityWriter with a mocked appender.")
	{
		appender := mockAppender{}
		writer := NewCityWriter("test", &appender)
		t.Log("When there are some cities to write.")
		{
			for _, city := range []*City{
				NewCity("Prague (1,350,000)"),
				NewCity("Brno (380,000)"),
				NewCity("Liberec (104,000)"),
			} {
				writer.Write(city)
			}
			writer.Flush()

			if appender.captured == expectedOutput {
				t.Log("Then writer writes cities as expected.")
			} else {
				t.Errorf("Then writer should've written cities in the expected format, but got this: %s",
					appender.captured)
			}
		}
	}
}

type mockAppender struct {
	captured string
}

func (this *mockAppender) Append(_ string, content string) {
	this.captured += content
}

const expectedOutput = `=== test ===
Prague (1,350,000)
Brno (380,000)
Liberec (104,000)
`
