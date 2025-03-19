package main


import (
   "fmt"
   "math"
)


func main() {
   oneProportionZTest()
}


func NormCdf(x float64, loc float64, scale float64) float64 {
   return 0.5 * (1 + math.Erf((x-loc)/(scale*math.Sqrt(2))))
}

func oneProportionZTest() {
   var sampleProportion float64
   var sampleSize float64
   var nullValue float64
   var significanceLevel float64
   var comparison string
   var isRandom bool
   var replacement bool
   var isSampleLessThanTenPercent bool

// Takes input for sample proportion
   sampleProportionAnswered := false
   for !sampleProportionAnswered {
       fmt.Println("Please enter the sample proportion (between 0 and 1).")
       fmt.Scan(&sampleProportion)
       if sampleProportion >= 0 && sampleProportion <= 1 {
           sampleProportionAnswered = true
       } else {
           fmt.Println("Please enter a vaid response.")
       }
   }

// Takes input for sample size
   sampleSizeAnswered := false
   for !sampleSizeAnswered {
       fmt.Println("Please enter the sample size (should be an integer).")
       fmt.Scan(&sampleSize)
       if sampleSize == math.Floor(sampleSize) {
           sampleSizeAnswered = true
       } else {
           fmt.Println("Please enter a valid response.")
       }
   }

// Takes input for Null Hypothesis Value
   nullValueAnswered := false
   for !nullValueAnswered {
       fmt.Println("Please enter the null value being compared (between 0 and 1).")
       fmt.Scan(&nullValue)
       if nullValue >= 0 && nullValue <= 1 {
           nullValueAnswered = true
       } else {
           fmt.Println("Please enter a valid response.")
       }
   }

// Takes input for comparison choice
   comparisonAnswered := false
   for !comparisonAnswered {
       fmt.Println("Please enter the alternative hypotheses (if it is that the population proportion is greater than null, enter 'G'; if it is that the population proportion is less than the null, enter 'L'; if it is that the population proportion is not equal to the null, enter 'U')")
       fmt.Scan(&comparison)
       if comparison == "G" || comparison == "L" || comparison == "U" {
           comparisonAnswered = true
       } else {
           fmt.Println("Please enter a valid response.")
       }
   }

// Takes input for significance level
   significanceLevelAnswered := false
   for !significanceLevelAnswered {
       fmt.Println("Please enter the significance level of the test (between 0 and 1).")
       fmt.Scan(&significanceLevel)
       if significanceLevel >= 0 && significanceLevel <= 1 {
           significanceLevelAnswered = true
       } else {
           fmt.Println("Please enter a valid response.")
       }
   }

// Takes input for whether or not sampling is random
   var actuallyIsRandom string
   isRandomAnswered := false
   for !isRandomAnswered {
       fmt.Println("Was the sample randomly sampled? Enter 'Yes' or 'No.'")
       fmt.Scan(&actuallyIsRandom)
       if actuallyIsRandom == "Yes" {
           isRandom = true
           isRandomAnswered = true
       } else if actuallyIsRandom == "No" {
           isRandom = false
           isRandomAnswered = true
       } else {
           fmt.Println("Please enter a valid response.")
       }
   }

// Takes input for whether or not sampling uses replacement
   var actuallyReplacement string
   replacementAnswered := false
   for !replacementAnswered {
       fmt.Println("Was the sample sampled with replacement. Enter 'Yes' or 'No.'")
       fmt.Scan(&actuallyReplacement)
       if actuallyReplacement == "Yes" {
           replacement = true
           replacementAnswered = true
       } else if actuallyReplacement == "No" {
           replacement = false
           replacementAnswered = true
       } else {
           fmt.Println("Please enter a valid response.")
       }
   }

// Takes input for whether or not sampling size is less than or equal to 10% of the population of interest
   var actuallyIsSampleLessThanTenPercent string
   isSampleLessThanTenPercentAnswered := false
   for !isSampleLessThanTenPercentAnswered {
       fmt.Println("Can it be reliably assumed that the given sample size is less than 10% of the size of the population of interest? Enter 'Yes' or 'No.'")
       fmt.Scan(&actuallyIsSampleLessThanTenPercent)
       if actuallyIsSampleLessThanTenPercent == "Yes" {
           isSampleLessThanTenPercent = true
           isSampleLessThanTenPercentAnswered = true
       } else if actuallyIsSampleLessThanTenPercent == "No" {
           isSampleLessThanTenPercent = false
           isSampleLessThanTenPercentAnswered = true
       } else {
           fmt.Println("Please enter a valid response.")
       }
   }

// Tests for conditions (random sampling, 10% Condition, Large Counts Rule) and calculates statistical probability value (p-value) using the NormCdf() math function
   if (comparison == "G" && sampleProportion > nullValue) || (comparison == "L" && sampleProportion < nullValue) || (comparison == "U" && sampleProportion != nullValue) {
       if isRandom && (replacement || isSampleLessThanTenPercent) && nullValue*sampleSize >= 10 && (1-nullValue)*sampleSize >= 10 {
           testStat := (sampleProportion - nullValue) / math.Sqrt((nullValue*(1-nullValue))/sampleSize)
           var pValue float64
           if sampleProportion <= nullValue {
               pValue = NormCdf(testStat, 0.00, 1.00)
           } else {
               pValue = 1 - NormCdf(testStat, 0.00, 1.00)
           }


           if comparison == "U" {
               pValue *= 2
           }

// Outputs z-test results
           if pValue <= significanceLevel {
               if comparison == "G" {
                   fmt.Printf("As the p-value of this test of about %.3f is less than the signficiance level of %.3f, the null hypothesis is rejected. There is convincing statistical evidence that the population proportion of the population of interest is greater than the null value.", pValue, significanceLevel)
               } else if comparison == "L" {
                   fmt.Printf("As the p-value of this test of about %.3f is less than the signficiance level of %.3f, the null hypothesis is rejected. There is convincing statistical evidence that the population proportion of the population of interest is less than the null value.", pValue, significanceLevel)
               } else if comparison == "U" {
                   fmt.Printf("As the p-value of this test of about %.3f is less than the signficiance level of %.3f, the null hypothesis is rejected. There is convincing statistical evidence that the population proportion of the population of interest is not equal to the null value.", pValue, significanceLevel)
               }
           } else {
               if comparison == "G" {
                   fmt.Printf("As the p-value of this test of about %.3f is greater than the signficiance level of %.3f, the null hypothesis fails to be rejected. There is not convincing statistical evidence that the population proportion of the population of interest is greater than the null value.", pValue, significanceLevel)
               } else if comparison == "L" {
                   fmt.Printf("As the p-value of this test of about %.3f is greater than the signficiance level of %.3f, the null hypothesis fails to be rejected. There is not convincing statistical evidence that the population proportion of the population of interest is less than the null value.", pValue, significanceLevel)
               } else if comparison == "U" {
                   fmt.Printf("As the p-value of this test of about %.3f is greater than the signficiance level of %.3f, the null hypothesis fails to be rejected. There is not convincing statistical evidence that the population proportion of the population of interest is not equal to the null value.", pValue, significanceLevel)
               }
           }
       } else {
           fmt.Println("Sorry. The statistical conditions necessary to perform a One Proportion Z-Test have not been met. Please make sure that the sample is randomly sampled, that the sample is sampled with replacement or that the it can be reliably assumed the sample size is less than 10% of the size of the population of interest, and that the sample size is sufficiently large (the sample size multiplied by the null value and multiplied by 1- the null value should be greater than or equal to 10).")
       }
   } else {
       fmt.Println("Sorry. Based on the sample statistic provided, there is not sufficient evidence to reject the null hypothesis.")
   }
}


