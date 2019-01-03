package test

import (
	"fmt"
	"github.com/clbanning/mxj"
	"testing"
)

func TestXmlParse(t *testing.T) {
	x := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!--
  ~ /* Copyright 2013-2014 the original author or authors.
  ~  *
  ~  * Licensed under the Apache License, Version 2.0 (the "License");
  ~  * you may not use this file except in compliance with the License.
  ~  * You may obtain a copy of the License at
  ~  *
  ~  *     http://www.apache.org/licenses/LICENSE-2.0
  ~  *
  ~  * Unless required by applicable law or agreed to in writing, software
  ~  * distributed under the License is distributed on an "AS IS" BASIS,
  ~  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  ~  * See the License for the specific language governing permissions and
  ~  * limitations under the License.
  ~  */
  -->

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
	m, _ := mxj.NewMapXml([]byte(x))
	fmt.Println(m)

}