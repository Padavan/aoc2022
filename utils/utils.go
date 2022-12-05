package utils;

import (
  "os"
  "log"
  "bufio"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  var lines []string
  // optionally, resize scanner's capacity for lines over 64K, see next example
  for scanner.Scan() {
      var line = scanner.Text()
      lines = append(lines, line);
      // fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  return lines
}

func IndexOf(element string, data []string) (int) {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1
}