web di golang kita perlu membuat routing, untuk menyediakan url.
di golang kita gunakan mux/server mux: `mux := http.NewServeMux()`,
lalu bisa membuat patter url menggunakan `mux.HandleFunc("/hello", hellowHandler)`.

closure adalah function yang disimpan dalam variable (hanya bisa di dalam func main?).

anonymous function adalah function yang didefinisikan tanpa nama, biasanya sebagai parameter atau sebuah closure.
