#!/bin/sh
cd ../tools
./playgame.py \
	-m maps/random_walk/random_walk_05p_02.map \
	"../go/MyBot" \
	"../go/v13" \
	"python sample_bots/python/HunterBot.py" \
	"python sample_bots/python/HunterBot.py" \
	"python sample_bots/python/HunterBot.py" \
	--fill \
	--log_dir game_logs --turns 399 \
	--food random \
	--player_seed 7 --verbose \
	--log_stderr \

# 	-m maps/random_walk/random_walk_02p_02.map \
# 	-m maps/random_walk/random_walk_02p_02.map \
