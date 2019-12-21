package main

import (
	"testing"
	. "github.com/stretchr/testify/assert"
	. "github.com/asuahsahua/advent2019/cmd/common"
)

func TestConfirmedAnswerPart1(t *testing.T) {
	robot := NewRobot()
	hull := NewHull()
	robot.RunPaintingProgram(hull)
	Equal(t, 2129, len(hull.Panels))
}

func TestConfirmedAnswerPart2(t *testing.T) {
	robot := NewRobot()
	hull := NewHull()
	hull.Panels[Point2D{0, 0}] = WHITE
	robot.RunPaintingProgram(hull)
	Equal(t, ExpectedPart2, hull.ToString())
}

var ExpectedPart2 = `######
  #  #
  #  #
   ## 
      
######
#  # #
#  # #
#    #
      
 #### 
#    #
#    #
 #  # 
      
######
   #  
 ## # 
#    #
      
######
  #  #
 ##  #
#  ## 
      
 #### 
#    #
# #  #
### # 
      
##   #
# #  #
#  # #
#   ##
      
######
#     
#     
#     
`