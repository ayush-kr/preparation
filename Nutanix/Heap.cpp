#include<bits/stdc++.h>

using namespace std;
#define MAX_SIZE 50

class Heap{
    private:
        int arr[MAX_SIZE];
        int size;
        int length = MAX_SIZE;
    public:
        Heap(){
            size = 0;
            length = MAX_SIZE;
        }

        void insertHeap(int val){
            if(size == length){
                cout<<"Cannot insert"<<endl;
                return;
            }
            arr[size] = val;
            size++;
            int i = size-1;
            while(i!=0 && arr[parent(i)]<arr[i]){
                int temp = arr[i];
                arr[i] = arr[parent(i)];
                arr[parent(i)] = temp;
                i=parent(i);
            }
        }

        void max_heapify(int ind){
            if(ind<=0 && ind>size)
                return;

            int l = left(ind);
            int r = right(ind);

            int largest = ind;
            if(l<=size-1 && arr[ind]<arr[l]){
                largest = l;
            }
            if(r<=size-1 && arr[largest]<arr[r]){
                largest = r;
            }

            if(largest!=ind){
                int temp = arr[largest];
                arr[largest] = arr[ind];
                arr[ind] = temp;
                max_heapify(largest);
            }
        }

        int extract_max(){
            if(size==0){
                cout<<"Cannot extract more"<<endl;
                return -1;
            }

            int max_ele = arr[0];
            arr[0] = arr[size-1];
            size--;

            max_heapify(0);
            return max_ele;
        }

        void print_heap(){
            if(size == 0){
                cout<<"Nothing in heap"<<endl;
                return;
            }
            cout<<"Heap elements:"<<endl;
            for(int i=0;i<size;i++)
                cout<<arr[i]<<" ";
            cout<<endl;
        }

        int parent(int i){
            return (i-1)/2;
        }
        int left(int i){
            return 2*i+1;
        }
        int right(int i){
            return 2*i+2 ;
        }
};

int main(){
    Heap heap;
    heap.print_heap();
    heap.insertHeap(5);
    heap.insertHeap(7);
    heap.insertHeap(3);
    heap.print_heap();
    cout<<"Current_max: "<<heap.extract_max()<<endl;
    cout<<"Current_max: "<<heap.extract_max()<<endl;
    heap.print_heap();
    cout<<"Current_max: "<<heap.extract_max()<<endl;
    heap.print_heap();
    cout<<"Current_max: "<<heap.extract_max()<<endl;
}