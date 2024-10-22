package bootcamp 

// MagicGrowthN returns all possible combinations of n different digits in ascending order.
func MagicGrowthN(n int) []string {
	if n < 1 || n > 10 {
		return []string{} // Return empty slice for invalid n
	}

	var result []string // Initialize the slice to store results

	// Use nested loops to generate combinations
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			if n == 2 {
				// Store the combination of two digits
				result = append(result, string('0'+i)+string('0'+j))
			} else {
				for k := j + 1; k <= 9; k++ {
					if n == 3 {
						// Store the combination of three digits
						result = append(result, string('0'+i)+string('0'+j)+string('0'+k))
					} else {
						for l := k + 1; l <= 9; l++ {
							if n == 4 {
								// Store the combination of four digits
								result = append(result, string('0'+i)+string('0'+j)+string('0'+k)+string('0'+l))
							} else {
								for m := l + 1; m <= 9; m++ {
									if n == 5 {
										// Store the combination of five digits
										result = append(result, string('0'+i)+string('0'+j)+string('0'+k)+string('0'+l)+string('0'+m))
									} else {
										for n1 := m + 1; n1 <= 9; n1++ {
											if n == 6 {
												// Store the combination of six digits
												result = append(result, string('0'+i)+string('0'+j)+string('0'+k)+string('0'+l)+string('0'+m)+string('0'+n1))
											} else {
												for n2 := n1 + 1; n2 <= 9; n2++ {
													if n == 7 {
														// Store the combination of seven digits
														result = append(result, string('0'+i)+string('0'+j)+string('0'+k)+string('0'+l)+string('0'+m)+string('0'+n1)+string('0'+n2))
													} else {
														for n3 := n2 + 1; n3 <= 9; n3++ {
															if n == 8 {
																// Store the combination of eight digits
																result = append(result, string('0'+i)+string('0'+j)+string('0'+k)+string('0'+l)+string('0'+m)+string('0'+n1)+string('0'+n2)+string('0'+n3))
															} else {
																for n4 := n3 + 1; n4 <= 9; n4++ {
																	if n == 9 {
																		// Store the combination of nine digits
																		result = append(result, string('0'+i)+string('0'+j)+string('0'+k)+string('0'+l)+string('0'+m)+string('0'+n1)+string('0'+n2)+string('0'+n3)+string('0'+n4))
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return result // Return the slice containing all combinations
}

// func main() {
// 	// Test the function with n=2
// 	result2 := MagicGrowthN(2)
// 	println("Combinations of 2 digits:")
// 	for _, val := range result2 {
// 		println(val) // Print each combination
// 	}
	
// 	// Test the function with n=3
// 	result3 := MagicGrowthN(3)
// 	println("Combinations of 3 digits:")
// 	for _, val := range result3 {
// 		println(val) // Print each combination
// 	}
	
// 	// Test for n=0 (invalid case)
// 	println("Output for n=0:")
// 	result0 := MagicGrowthN(0) // Expected []
// 	for _, val := range result0 {
// 		println(val) // Should not print anything
// 	}

// 	// Test for n=11 (invalid case)
// 	println("Output for n=11:")
// 	result11 := MagicGrowthN(11) // Expected []
// 	for _, val := range result11 {
// 		println(val) // Should not print anything
// 	}
// }
