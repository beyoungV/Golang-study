#include<iostream>
#include<opencv2/opencv.hpp>
using namespace std;
using namespace cv;
int main(int argc,const char** argv)
    {
            cout << "Load image from E:\\BFV.jpg" << endl;
            Mat src = imread("E:\\BFV.jpg");
            imshow("BFV",src);
            waitKey(0);
            return 0;
    }



//g++ .\main.cpp -o main.exe -I E:\opencv\include -L E:\opencv\x64\mingw\lib -l opencv_img_hash345 -l opencv_world345

