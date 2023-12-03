#!/usr/bin/env ruby

depth = 0
horizontal = 0
aim = 0

commands = File.readlines("./data/day-2.txt").map do |command|
  { action: command.split.first, value: command.split.last }
end

commands.each do |command|
  value = command[:value].to_i
  action = command[:action]

  if action == 'forward'
    horizontal += value
    depth = depth + aim * value
  end

  aim -= value if action == 'up'
  aim += value if action == 'down'
end

puts horizontal * depth
