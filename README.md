# one-proportion-z-test-program
A Golang program designed to conduct statistical Z-Tests

Overview: This program is designed to conduct One Proportion Statistical Z-Tests. This program was written in the
Golang programming language and utilizes statistical laws and formulas to provide users such as researchers
and data specialists using sample proportions a way to evaluate whether or not there is convincing statistical evidence to prove that 
a population proportion differs from a null hypothesis. This program can be highly relevant in a variety of areas of research,
such as researching consumer and market preferences and interests for economic applications, socioeconomic population
characteristics, and public opinion.

Functionality: This program takes necessary inputs from users to conduct a proper One Proportion Statistical Z-Test. This includes
sample proportion for a proportion of interest, sample size, Null Hypothesis value, the comparison to test for (greater than, less than, 
or unequal to the Null Hypothesis value) significance level, whether or not random sampling was used, whether or
not replacement was used in sampling, and whether or not less than or equal to 10% of the population of interest was used for
the sampling. The program will then check for the necessary conditions of testing, including random sampling, the 10% Condition
for sampling without replacement, and the Large Counts Rule. If the conditions for testing are made, the program will make statistical
calculations with the use of functions such as NormCDF to calculate the probability of the user's sampling proportion compared to the
Null Hypothesis Value in the population of interest and determine whether or not the sample proportion gives convincing statistical evidence
that the population proportion of interest is greater than, less than, or unequal to the Null Hypothesis value based on the given
significance level.

Conclusion: By making this model, I was able to combine my knowledge and skills of both the Golang programming language as well as
Statistics towards a real-life applications. I hope to build upon my projects to develop further computational statistical and
information-processing model in the future in areas such as economic and financial research.
