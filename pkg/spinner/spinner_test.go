package spinner

import (
	"errors"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/kaduartur/go-cli-spinner/pkg/template"
)

func TestSpinner_Default_Config(t *testing.T) {
	title := "Loading..."
	spinner := StartNew(title)
	spinner.Stop()

	if spinner.Title != title {
		t.Errorf("The title expected %s and got %s", title, spinner.Title)
	}

	if spinner.FrameRate != DefaultFrameRate {
		t.Errorf("The FrameRate expected %v and got %v", DefaultFrameRate, spinner.FrameRate)
	}

	if !reflect.DeepEqual(spinner.Template, template.Default) {
		t.Errorf("The Template expected %v and got %v", template.Default, spinner.Template)
	}

	if !spinner.NoTty {
		t.Errorf("The NoTty expected %v and got %v", true, spinner.NoTty)
	}
}

func TestSpinner_custom_config(t *testing.T) {
	title := "Loading..."
	speed := time.Millisecond + 100
	spinner := New(title)
	spinner.NoTty = false
	spinner.SetTemplate(template.Arrow)
	spinner.SetSpeed(speed)
	spinner.Start()
	spinner.NoTty = false
	time.Sleep(time.Second * 2)
	spinner.Stop()

	if spinner.Title != title {
		t.Errorf("The title expected %s and got %s", title, spinner.Title)
	}

	if spinner.FrameRate != speed {
		t.Errorf("The FrameRate expected %v and got %v", speed, spinner.FrameRate)
	}

	if !reflect.DeepEqual(spinner.Template, template.Arrow) {
		t.Errorf("The Template expected %v and got %v", template.Arrow, spinner.Template)
	}

	if spinner.NoTty {
		t.Errorf("The NoTty expected %v and got %v", true, spinner.NoTty)
	}
}

func TestSpinner_custom_output(t *testing.T) {
	title := "Loading..."
	spinner := New(title)
	spinner.Output = os.Stdout
	spinner.Start()
	time.Sleep(time.Second * 2)
	spinner.Stop()
	spinner.Success("OK!")
	spinner.Error(errors.New("error"))

	if spinner.Title != title {
		t.Errorf("The title expected %s and got %s", title, spinner.Title)
	}

	if spinner.FrameRate != DefaultFrameRate {
		t.Errorf("The FrameRate expected %v and got %v", DefaultFrameRate, spinner.FrameRate)
	}

	if !reflect.DeepEqual(spinner.Template, template.Default) {
		t.Errorf("The Template expected %v and got %v", template.Arrow, spinner.Template)
	}

	if !spinner.NoTty {
		t.Errorf("The NoTty expected %v and got %v", true, spinner.NoTty)
	}
}