package config

import "time"

// With an input signal "of" 62.5 ms (1 / 2^4)
const INPUT_SIGNAL_FREQUENCY = 62500 * time.Microsecond

// We can "divide" it by 4 to get our signal "of" 1 second
const FLIP_FLOPS_COUNT = 4

// Log source input (signal) and intermediate flip-flops outputs
const DEBUG_LOGS = false
