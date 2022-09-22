package restic

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// print is used to output stdout and stderr in real time.
func print(stdout, stderr io.ReadCloser, done chan struct{}) {
	stopCh := make(chan struct{}, 2)
	defer stdout.Close()
	defer stderr.Close()
	//errCh := make(chan error, 2)

	// A goroutine that outputs stdout in real time.
	go func() {
		defer func() {
			stopCh <- struct{}{}
		}()
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Printf("%s", scanner.Text())
		}
		err := scanner.Err()
		// if stdout already closed, stop the goroutine.
		if errors.Is(err, os.ErrClosed) {
			return
		}
		if err != nil {
			fmt.Println("scanner output stdout error:", err)
			return
		}
	}()

	// A goroutine that outputs stderr in real time.
	go func() {
		defer func() {
			stopCh <- struct{}{}
		}()
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Printf("%s", scanner.Text())
		}
		err := scanner.Err()
		// if stderr already closed, stop the goroutine.
		if errors.Is(err, os.ErrClosed) {
			return
		}
		if err != nil {
			fmt.Println("scanner output stderr error:", err)
			return
		}
	}()

	<-stopCh
	<-stopCh
	done <- struct{}{}
}

// concatFlags will concatenates all restic command flags into string.
// "omitempty" tag works on "int", "int32", "int64" , []string and "map[string]string" types.
func concatFlags(flags interface{}) string {
	var (
		t reflect.Type
		v reflect.Value
		s string
	)
	if reflect.ValueOf(flags).Kind() == reflect.Pointer {
		t = reflect.TypeOf(flags).Elem()
		v = reflect.ValueOf(flags).Elem()
	} else {
		t = reflect.TypeOf(flags)
		v = reflect.ValueOf(flags)
	}

	for i := 0; i < v.NumField(); i++ {
		// PkgPath is the package path that qualified a lows case (unexported)
		// field name. It is empty for upper case (exported) field names.
		// Skip unexported fields
		if len(t.Field(i).PkgPath) != 0 {
			continue
		}
		knd := v.Field(i).Kind()
		typ := v.Field(i).Type()
		val := v.Field(i).Interface()
		nam := t.Field(i).Name
		//nam := v.Type().Field(i).Name
		tag := t.Field(i).Tag.Get("json")

		if !v.IsValid() {
			continue
		}
		if v.IsZero() {
			continue
		}
		// If json tag have multiple values, the frist value separated by "," is
		// restic flag name.
		flagName := strings.Split(tag, ",")[0]
		omitempty := strings.Contains(tag, "omitempty")

		_ = knd
		_ = typ
		_ = val
		_ = nam
		_ = tag

		switch typ.String() {
		case "string":
			flagValue, ok := val.(string)
			if ok {
				if len(flagValue) == 0 {
					continue
				}
				s = s + " " + flagName + "=" + flagValue
				s = strings.TrimSpace(s)
			}
		case "[]string":
			flagValue, ok := val.([]string)
			if ok {
				if flagValue == nil {
					continue
				}
				if len(flagValue) == 0 && omitempty {
					continue
				}
				s = s + " " + flagName + "=" + strings.Join(flagValue, ",")
				s = strings.TrimSpace(s)
			}
		case "int":
			flagValue, ok := val.(int)
			if ok {
				if flagValue == 0 && omitempty {
					continue
				}
				s = s + " " + flagName + "=" + strconv.Itoa(flagValue)
				s = strings.TrimSpace(s)
			}
		case "int32":
			flagValue, ok := val.(int32)
			if ok {
				if flagValue == 0 && omitempty {
					continue
				}
				s = s + " " + flagName + "=" + strconv.FormatInt(int64(flagValue), 10)
				s = strings.TrimSpace(s)
			}
		case "int64":
			flagValue, ok := val.(int64)
			if ok {
				if flagValue == 0 && omitempty {
					continue
				}
				s = s + " " + flagName + "=" + strconv.FormatInt(flagValue, 10)
				s = strings.TrimSpace(s)
			}
		case "bool":
			flagValue, ok := val.(bool)
			if ok {
				if flagValue == false {
					continue
				}
				s = s + " " + flagName
				s = strings.TrimSpace(s)
			}
		case "map[string]string":
			flagValue, ok := val.(map[string]string)
			if ok {
				if flagValue == nil {
					continue
				}
				if len(flagValue) == 0 && omitempty {
					continue
				}
				var ts string
				for key, val := range flagValue {
					ts = ts + key + "=" + val + ","
				}
				ts = strings.TrimSuffix(ts, ",")
				s = s + " " + flagName + "=" + ts
				s = strings.TrimSpace(s)
			}
		}
	}

	return s
}

// envMapToSlice convert a map containing environment variables to silce.
func envMapToSlice(envMap map[string]string) []string {
	if envMap == nil {
		return nil
	}

	var envSlice []string
	for k, v := range envMap {
		envSlice = append(envSlice, k+"="+v)
	}
	return envSlice
}

// deduplicateStrSlice removes duplicates in string slice.
// ref: https://stackoverflow.com/questions/66643946/how-to-remove-duplicates-strings-or-int-from-slice-in-go
func deduplicateStrSlice(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var newSlice []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

// deduplicateIntSlice reomves duplicates in int slice.
// ref: https://stackoverflow.com/questions/66643946/how-to-remove-duplicates-strings-or-int-from-slice-in-go
func deduplicateIntSlice(intSlice []int) []int {
	allKeys := make(map[int]bool)
	var newSlice []int
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}
