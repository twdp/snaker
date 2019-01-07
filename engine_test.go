package snaker

import "testing"

var xml = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>

<process displayName="借款申请流程" instanceUrl="/snaker/flow/all" name="borrow">
    <start displayName="start1" layout="42,118,-1,-1" name="start1">
        <transition g="" name="transition1" offset="0,0" to="apply"/>
    </start>
    <end displayName="end1" layout="479,118,-1,-1" name="end1"/>
    <task assignee="apply.operator" autoExecute="Y" displayName="借款申请" form="/flow/borrow/apply" layout="126,116,-1,-1" name="apply" performType="ANY" taskType="Major">
        <transition g="" name="transition2" offset="0,0" to="approval"/>
    </task>
    <task assignee="approval.operator" autoExecute="Y" displayName="审批" form="/snaker/flow/approval" layout="252,116,-1,-1" name="approval" performType="ANY" taskType="Major">
     <transition g="" name="transition3" offset="0,0" to="decision1"/>
    </task>
    <decision displayName="decision1" expr="#result" layout="384,118,-1,-1" name="decision1">
        <transition displayName="同意" g="" name="agree" offset="0,0" to="end1"/>
        <transition displayName="不同意" g="408,68;172,68" name="disagree" offset="0,0" to="apply"/>
    </decision>
</process>
`

func TestSnakerEngineImpl_StartInstanceById(t *testing.T) {

}