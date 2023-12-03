#!/usr/bin/env ruby

numbers = File.readlines("./data/day-1.txt")
result = 0

numbers.each_with_index do |number, index|
  next unless index > 0

  result += 1 if number.to_i > numbers[index.to_i - 1].to_i
end

puts result
