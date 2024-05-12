package scripts_test

import (
	"errors"
	"fmt"
	"main/internal/model"
	"main/internal/scripts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand_RunScript(t *testing.T) {

	testCases := []struct {
		name    string
		command interface{}
		result  string
		errors  error
	}{
		{
			name: "valid",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			result: "Hello, World Test!!!\n",
			errors: nil,
		},
		{
			name: "empty script",
			command: model.Command{
				Script: "",
			},
			result: "",
			errors: errors.New("invalid script"),
		},
		{
			name: "invalid script",
			command: model.Command{
				Script: "#!/bin/bash\nech \"Hello, World Test!!!\"",
			},
			result: "",
			errors: errors.New("invalid script"),
		},
		{
			name: "test script 1",
			command: model.Command{
				Script: `#!/bin/bash
				url="http://example.com"
				http_status=$(curl -s -o /dev/null -I -w "%{http_code}" "$url")
				if [ "$http_status" -eq 200 ]; then
					echo "$url доступен, код ответа: $http_status"
				else
					echo "$url недоступен, код ответа: $http_status"
				fi`,
			},
			result: "http://example.com доступен, код ответа: 200\n",
			errors: nil,
		},
		{
			name: "test script_ check site",
			command: model.Command{
				Script: `#!/bin/bash
				url="http://ex00mpl.co"
				http_status=$(curl -s -o /dev/null -I -w "%{http_code}" "$url")
				if [ "$http_status" -eq 200 ]; then
					echo "$url доступен, код ответа: $http_status"
				else
					echo "$url недоступен, код ответа: 404"
				fi`,
			},
			result: "http://ex00mpl.co недоступен, код ответа: 404\n",
			errors: nil,
		},
		{
			name: "test script find file",
			command: model.Command{
				Script: `#!/bin/bash
				search_term="testfile*"

				echo "Поиск файлов, содержащих '$search_term':"
				find . -type f -name "*$search_term*"`,
			},
			result: "Поиск файлов, содержащих 'testfile*':\n./testfile.txt\n",
			errors: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := scripts.Run(tc.command.(model.Command).Script)
			if err != nil {
				err = fmt.Errorf("invalid script")
			}

			assert.Equal(t, tc.errors, err)
			assert.Equal(t, tc.result, string(result))

		})

	}
}

/*
Вот несколько простых bash-скриптов:
Скрипт для вывода текущей даты и времени:
bash
#!/bin/bash
echo "Текущая дата и время:"
date

Скрипт для вывода списка файлов в текущей директории:
bash
#!/bin/bash
echo "Содержимое текущей директории:"
ls -l

Скрипт для создания резервной копии файла:
bash
#!/bin/bash
file="file_to_backup.txt"
backup_file="$file.backup"

cp "$file" "$backup_file"
echo "Резервная копия файла $file создана: $backup_file"

Скрипт для поиска файлов по имени:
bash
#!/bin/bash
search_term="$1"

echo "Поиск файлов, содержащих '$search_term':"
find . -type f -name "*$search_term*"

Скрипт для отправки email:
bash
#!/bin/bash
to="recipient@example.com"
subject="Тема письма"
body="Содержимое письма"

echo "$body" | mail -s "$subject" "$to"

Скрипт для создания директории и перехода в нее:
bash
#!/bin/bash
dir_name="new_directory"

mkdir "$dir_name"
cd "$dir_name"
echo "Создана и открыта директория: $dir_name"


*/
