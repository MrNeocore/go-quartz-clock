# Golang Quartz Clock demo

This project implements a "Quartz Clock" similar to a real digital watch / clock.

## Principle
A high frequency (a power of 2) input source signal is used to generate a 1 second output signal - serving as the "ticking" input for the clock.

Frequency dividers (also called flip-flops) are implemented as goroutine and share an input and output channel with the previous and next flip-flop (or are used as the source signal / output signal).

</br>

## Illustration

</br>

<a href="https://www.electronics-tutorials.ws/counter/count_1.html">
    <img src="frequency-divider.png" width="400">
</a>

</br>


## Usage
``` bash
go run main.go -durationSeconds 3
```

### Output example

```
Starting
=> 2022-02-13 23:37:30.488941189 +0100 CET m=+1.001053330
=> 2022-02-13 23:37:31.488604518 +0100 CET m=+2.000716655
=> 2022-02-13 23:37:32.488471204 +0100 CET m=+3.000583350
Exiting
```

## Note
Obviously, this project has no use beside explanation / training.

An equivalent, more performant and accurate implementation is simply :

``` golang
time.Ticker(1 * time.Second)
```

