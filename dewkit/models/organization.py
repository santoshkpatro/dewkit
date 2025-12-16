from django.db import models

from dewkit.models.shared import BaseModel


class Organization(BaseModel):
    name = models.CharField(max_length=128)
    website = models.URLField(blank=True, null=True)

    class Meta:
        db_table = "organizations"
