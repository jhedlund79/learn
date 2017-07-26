<?php
	/* 
		TO-DO:
	 	1. CLEAN INPUTS BEFORE DATABASE QUERY
	 	2. HANDLE ERRORS
	 	3. QUERY DATABASE AND RETURN RESULTS
	 */
	$user_id = $_POST['user_id'];
?>

<div class="modal-dialog">
    <div class="modal-content">
        <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h4 class="modal-title">User ID: <?php echo $user_id; ?> Details</h4>
        </div><!-- .modal-header -->
        <div class="modal-body">
            <p>One fine body&hellip;</p>
        </div><!-- .modal-body -->
        <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
        </div><!-- .modal-content -->
    </div><!-- .modal-content -->
</div><!-- .modal-dialog -->