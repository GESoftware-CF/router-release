#!bin/bash

# tee_output_to_sys_log
#
# When syslog_utils.sh is loaded, this sends stdout and stderr to /var/vcap/sys/log.
function tee_output_to_sys_log() {
  declare log_dir="$1"
  declare log_name="$2"

  exec > >(tee -a >(logger -p user.info -t "vcap.${log_name}.stdout") | prepend_datetime >>"${log_dir}/${log_name}.log")
  exec 2> >(tee -a >(logger -p user.error -t "vcap.${log_name}.stderr") | prepend_datetime >>"${log_dir}/${log_name}.err.log")
}

function prepend_datetime() {
  perl -ne 'BEGIN { use Time::HiRes "time"; use POSIX "strftime"; STDOUT->autoflush(1) }; my $t = time; my $fsec = sprintf ".%06d", ($t-int($t))*1000000; my $time = strftime("%Y-%m-%dT%H:%M:%S".$fsec."%z", localtime $t); print("$time $_")'
}
