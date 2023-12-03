#!/usr/bin/env ruby

$numbers = File.readlines("./data/day-1.txt").map(&:to_i)
$result = 0

def slide_ze_windows
  first_window = $numbers[0..2]
  second_window = $numbers[1..3]

  $result += 1 if second_window.sum > first_window.sum
  $numbers.shift

  slide_ze_windows if $numbers.length > 3
end

slide_ze_windows

puts $result
