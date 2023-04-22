package main

import (
  "context"
  "fmt"
  "os"
  "os/signal"
  "sync"
  "github.com/go-telegram/bot"
  "github.com/go-telegram/bot/models"
  "github.com/amantaiaitolkyn/GO/Assigment4/unsplash"
)

// const (
//   unsplashAPIBaseURL = "https://api.unsplash.com"
//   unsplashRandomPath = "/photos/random"
// )

// type unsplashResponse struct {
//   URLs struct {
//     Regular string `json:"regular"`
//   } `json:"urls"`
// }

func main() {
  ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
  defer cancel()

  wg := &sync.WaitGroup{}
  mutex := &sync.Mutex{}
  counter := 0

  opts := []bot.Option{
    bot.WithDefaultHandler(handler),
  }

  b, err := bot.New("6286804537:AAEQorf6DfCww5szsJ2gVrMRTFl6WL7Y9O0", opts...)
  if err != nil {
    panic(err)
  }
  b.RegisterHandler(bot.HandlerTypeMessageText, "/img", bot.MatchTypeExact, handler)

  b.Start(ctx)

  for i := 0; i < 10; i++ {
    wg.Add(1)
    go CountWaitGroup(wg, &counter)
  }

  for i := 0; i < 10; i++ {
    wg.Add(1)
    go CountWithMutex(wg, mutex, &counter)
  }

  counterCh := make(chan int)
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go CountWithChanel(wg, counterCh)
  }

  go func() {
    for {
      select {
      case c := <-counterCh:
        mutex.Lock()
        counter += c
        mutex.Unlock()
        wg.Done()
      case <-ctx.Done():


        
        return
      }
    }
  }()

  defer func() {
    fmt.Println("Counter value:", counter)
  }()

  wg.Wait()
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
  if update.Message == nil {
    return
  }

  imageURL, err := b.GetRandomUnsplashImageURL()
  if err != nil {
    b.SendMessage(ctx, &bot.SendMessageParams{
      ChatID: update.Message.Chat.ID,
      Text:   "DIDN'T GET IT",
    })
    return
  }

  b.SendMessage(ctx, &bot.SendMessageParams{
    ChatID: update.Message.Chat.ID,
    Text:   imageURL,
  })
}

// func getRandomUnsplashImageURL() (string, error) {
//   client := &http.Client{}
//   req, err := http.NewRequest(http.MethodGet, unsplashAPIBaseURL+unsplashRandomPath, nil)
//   if err != nil {
//     return "", err
//   }
//   req.Header.Add("Authorization", "Client-ID APhyMWrOZfZbKgVkgLk9QNU-CgVXmtX43WqD88BBG8M")

//   resp, err := client.Do(req)
//   if err != nil {
//     return "", err
//   }
//   defer resp.Body.Close()

//   var unsplashResp unsplashResponse
//   err = json.NewDecoder(resp.Body).Decode(&unsplashResp)
//   if err != nil {
//     return "", err
//   }
//   return unsplashResp.URLs.Regular, nil
// }

func CountWaitGroup(wg *sync.WaitGroup, counter *int) {
  defer wg.Done()
  *counter++
}

func CountWithMutex(wg *sync.WaitGroup, mutex *sync.Mutex, counter *int) {
  defer wg.Done()
  mutex.Lock()
  *counter++
  mutex.Unlock()
}

func CountWithChanel(wg *sync.WaitGroup, counterCh chan<- int) {
  defer wg.Done()
  counterCh <- 1
}