#!/bin/sh
T=../tools
make && $T/playgame.py \
	--engine_seed 42 \
	--player_seed 42 \
	--food none \
	--end_wait=0.25 \
	--verbose \
	--log_dir $T/game_logs \
	--turns 30 \
	--map_file $T/submission_test/test.map \
	 --nolaunch \
	-e \
	--strict \
	--capture_errors \
	"../go/MyBot" \
	"python $T/submission_test/TestBot.py" \
