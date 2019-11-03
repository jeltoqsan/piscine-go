package piscine

import "github.com/01-edu/z01"

// package main

// import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n > 0 && n < 10 {
		for i := '0'; i <= '9'; i++ {
			if n == 1 {
				z01.PrintRune(i)
				if i != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
			if n > 1 {

				for j := i + 1; j <= '9'; j++ {
					if n == 2 {
						z01.PrintRune(i)
						z01.PrintRune(j)
						if i != '8' || j != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
					if n > 2 {
						for k := j + 1; k <= '9'; k++ {
							if n == 3 {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								if i != '7' || j != '8' || k != '9' {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
							if n > 3 {
								for l := k + 1; l <= '9'; l++ {
									if n == 4 {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(l)
										if i != '6' || j != '7' || k != '8' || l != '9' {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
									if n > 4 {
										for m := l + 1; m <= '9'; m++ {
											if n == 5 {
												z01.PrintRune(i)
												z01.PrintRune(j)
												z01.PrintRune(k)
												z01.PrintRune(l)
												z01.PrintRune(m)
												if i != '5' || j != '6' || k != '7' || l != '8' || m != '9' {
													z01.PrintRune(',')
													z01.PrintRune(' ')
												}
											}
											if n > 5 {
												for o := m + 1; o <= '9'; o++ {
													if n == 6 {
														z01.PrintRune(i)
														z01.PrintRune(j)
														z01.PrintRune(k)
														z01.PrintRune(l)
														z01.PrintRune(m)
														z01.PrintRune(o)
														if i != '4' || j != '5' || k != '6' || l != '7' || m != '8' || o != '9' {
															z01.PrintRune(',')
															z01.PrintRune(' ')
														}
													}
													if n > 6 { //-----7----
														for p := o + 1; p <= '9'; p++ {
															if n == 7 {
																z01.PrintRune(i)
																z01.PrintRune(j)
																z01.PrintRune(k)
																z01.PrintRune(l)
																z01.PrintRune(m)
																z01.PrintRune(o)
																z01.PrintRune(p)
																if i != '3' || j != '4' || k != '5' || l != '6' || m != '7' || o != '8' || p != '9' {
																	z01.PrintRune(',')
																	z01.PrintRune(' ')
																}
															}
															if n > 7 {
																for q := p + 1; q <= '9'; q++ {
																	if n == 8 {
																		z01.PrintRune(i)
																		z01.PrintRune(j)
																		z01.PrintRune(k)
																		z01.PrintRune(l)
																		z01.PrintRune(m)
																		z01.PrintRune(o)
																		z01.PrintRune(p)
																		z01.PrintRune(q)
																		if i != '2' || j != '3' || k != '4' || l != '5' || m != '6' || o != '7' || p != '8' || q != '9' {
																			z01.PrintRune(',')
																			z01.PrintRune(' ')
																		}
																	}
																	if n > 8 {
																		for r := q + 1; r <= '9'; r++ {
																			if n == 9 {
																				z01.PrintRune(i)
																				z01.PrintRune(j)
																				z01.PrintRune(k)
																				z01.PrintRune(l)
																				z01.PrintRune(m)
																				z01.PrintRune(o)
																				z01.PrintRune(p)
																				z01.PrintRune(q)
																				z01.PrintRune(r)
																				if i != '1' || j != '2' || k != '3' || l != '4' || m != '5' || o != '6' || p != '7' || q != '8' || r != '9' {
																					z01.PrintRune(',')
																					z01.PrintRune(' ')
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
			}
		}
	}
	z01.PrintRune('\n')
}

// func main() {
// 	PrintCombN(1)
// 	PrintCombN(2)
// 	PrintCombN(9)
// }
