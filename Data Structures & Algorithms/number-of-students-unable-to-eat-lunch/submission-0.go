func countStudents(students []int, sandwiches []int) int {
    q := make([]int,len(students))
	copy(q, students)

	for _, sandwich := range sandwiches {
		cnt := 0
		for cnt < len(q) && q[0] != sandwich {
			q = append(q[1:], q[0])
			cnt++
		}
		// finished iterating 

		if q[0] == sandwich {
			q = q[1:]
		} else {
			break
		}
	}

	return len(q)
}