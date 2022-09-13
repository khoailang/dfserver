package data

import (
	"testing"
)

func TestParse(t *testing.T) {
	argslist := []string{
		"-h", "--help", "--tokenize", "-t", "--height", "-H", "--width", "-W",
		"--cfg_scale", "-C", "--number", "-n", "--separate-images", "-i", "--grid", "-g",
		"--sampler", "-A", "--steps", "-s", "--seed", "-S", "--prior", "-p"}

	cmd0 := "!dream half body portrait drawing of anime muscular horse, epic pose, pen and ink, intricate line drawings, by craig mullins, ruan jia, kentaro miura, greg rutkowski, loundraw and dan mumford"
	cmd0_cmd := "!dream half body portrait drawing of anime muscular horse, epic pose, pen and ink, intricate line drawings, by craig mullins, ruan jia, kentaro miura, greg rutkowski, loundraw and dan mumford"
	cmd0_args := []string{}
	args := ArgsParse(cmd0, argslist)
	verifyargs(t, args, cmd0_cmd, cmd0_args)

	cmd1 := "!dream half body portrait drawing of anime muscular horse, epic pose, pen and ink, intricate line drawings, by craig mullins, ruan jia, kentaro miura, greg rutkowski, loundraw and dan mumford -abc -W 768 -n 4 --seed 9958342083"
	cmd1_cmd := "!dream half body portrait drawing of anime muscular horse, epic pose, pen and ink, intricate line drawings, by craig mullins, ruan jia, kentaro miura, greg rutkowski, loundraw and dan mumford -abc"
	cmd1_args := []string{"-W 768", "-n 4", "--seed 9958342083"}
	args = ArgsParse(cmd1, argslist)
	verifyargs(t, args, cmd1_cmd, cmd1_args)

	cmd2 := "!dream half body portrait drawing of anime muscular horse, epic pose, pen and ink, intricate line drawings, by craig mullins, ruan jia, kentaro miura, greg rutkowski, loundraw and dan -abc mumford"
	cmd2_cmd := "!dream half body portrait drawing of anime muscular horse, epic pose, pen and ink, intricate line drawings, by craig mullins, ruan jia, kentaro miura, greg rutkowski, loundraw and dan -abc mumford"
	cmd2_args := []string{}
	args = ArgsParse(cmd2, argslist)
	verifyargs(t, args, cmd2_cmd, cmd2_args)

	cmd3 := "!dream half body portrait drawing of anime muscular horse, epic pose, pen and ink, intricate line drawings, by craig mullins, ruan jia, kentaro miura, greg rutkowski, loundraw and dan mumford --seed 9958342083"
	cmd3_cmd := "!dream half body portrait drawing of anime muscular horse, epic pose, pen and ink, intricate line drawings, by craig mullins, ruan jia, kentaro miura, greg rutkowski, loundraw and dan mumford"
	cmd3_args := []string{"--seed 9958342083"}
	args = ArgsParse(cmd3, argslist)
	verifyargs(t, args, cmd3_cmd, cmd3_args)

	cmd4 := "!dream half body portrait drawing of anime muscular horse, epic pose, pen and ink, intricate line drawings, by craig mullins, ruan jia, kentaro miura, greg rutkowski, loundraw and dan mumford -n 4 "
	cmd4_cmd := "!dream half body portrait drawing of anime muscular horse, epic pose, pen and ink, intricate line drawings, by craig mullins, ruan jia, kentaro miura, greg rutkowski, loundraw and dan mumford"
	cmd4_args := []string{"-n 4"}
	args = ArgsParse(cmd4, argslist)
	verifyargs(t, args, cmd4_cmd, cmd4_args)

	cmd5 := "!dream half -n 4 -U --face"
	cmd5_cmd := "!dream half"
	cmd5_args := []string{"-n 4", "-U", "--face"}
	args = ArgsParse(cmd5, argslist)
	verifyargs(t, args, cmd5_cmd, cmd5_args)
}

func verifyargs(t *testing.T, args *CommandArgs, rcmd string, rargs []string) {
	if args.Cmd != rcmd {
		t.Errorf("parse cmd err, \nexpect: (%d) %s\nresult: (%d) %s", len(rcmd), rcmd, len(args.Cmd), args.Cmd)
	}
	argok := true
	if len(args.Args) != len(rargs) {
		argok = false
	}
	for i, v := range args.Args {
		if v != rargs[i] {
			argok = false
		}
	}
	if argok == false {
		t.Errorf("parse args err, expect: %s ,result: %s ", rargs, args.Args)

	}
}
