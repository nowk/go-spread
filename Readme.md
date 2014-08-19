# go-spread

Spread jobs accross a set number of workers

## Example

    type Job struct {
      ID int
    }

    func (j Job) Do() error {
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
      log.Print("done", j.ID)
      return nil
    }

    func main() {
      a := Job{1}
      b := Job{2}
      c := Job{3}
      d := Job{4}

      errs := Spread(2, a, b, c, d)
      for err := range errs {
        log.Printf("error: %s", err)
      }
    }

### License

MIT
