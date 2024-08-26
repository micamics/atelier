from django.shortcuts import render
from django.http import HttpResponse
from .models import Movie

# Create your views here.


def index(request):
    movies = Movie.objects.all()
    output = ", ".join([m.title for m in movies])
    return HttpResponse(output)
