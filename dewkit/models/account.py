from django.db import models
from django.contrib.auth.models import AbstractBaseUser, BaseUserManager

from dewkit.models.shared import BaseModel


class Account(BaseModel, AbstractBaseUser):
    email = models.EmailField(unique=True)
    # password -> Already part of Abstract Class of django
    verified_at = models.DateTimeField(blank=True, null=True)

    class Meta:
        db_table = "accounts"
