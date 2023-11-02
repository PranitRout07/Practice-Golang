Here after executing kubectl port-forward service/my-service 8501:8501 , when there is a scale down in the number of pods due to less it throws some error
and the command stops executing. To solve this error i have written this golang program which will automatically re run the command when the error occurs.
