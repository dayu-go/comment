package store

const (
	_incrCmtIdxCntSQL   = "UPDATE comment_index_1 SET count=count+1,root_count=root_count+1,all_count=all_acount+1,update_time=? WHERE obj_id=? AND obj_type=?"
	_incrSubFloorCntSQL = "UPDATE reply_subject_1 SET count=count+1,update_time=? WHERE obj_id=? AND obj_type=?"
	_incrSubRootCntSQL  = "UPDATE reply_subject_1 SET root_count=root_count+1,update_time=? WHERE obj_id=? AND obj_type=?"
	_incrSubApplyCntSQL = "UPDATE reply_subject_1 SET all_count=all_count+1,update_time=? WHERE obj_id=? AND obj_type=?"
)
