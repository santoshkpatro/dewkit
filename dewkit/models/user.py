from django.db import models
from django.db import transaction

from dewkit.models.shared import BaseModel
from dewkit.models.account import Account
from dewkit.models.organization import Organization


class User(BaseModel):
    class Role(models.CharField):
        OWNER = ("owner", "Owner")
        ADMIN = ("admin", "Admin")
        STAFF = ("staff", "Staff")

    user_id = models.CharField(max_length=12, unique=True, blank=True)
    organization = models.ForeignKey(
        "Organization",
        on_delete=models.CASCADE,
        related_name="users",
        blank=True,
        null=True,
    )
    account = models.ForeignKey(
        "Account", on_delete=models.PROTECT, related_name="users", blank=True, null=True
    )

    # [Email] Redundant field as email is already present in account! (But needed for faster lookups)
    email = models.EmailField()
    full_name = models.CharField()
    is_superuser = models.BooleanField(default=False)
    role = models.CharField(max_length=16, choices=Role.choices, default=Role.STAFF)

    USERNAME_FIELD = "user_id"
    REQUIRED_FIELDS = ["email", "full_name"]

    class Meta:
        db_table = "users"

    @classmethod
    @transaction.atomic
    def create_account(cls, email, full_name, password, organization_name):
        # 1. Create an account
        new_account = Account(email=email)
        new_account.set_password(password)
        new_account.save()

        # 2. Create an organization
        new_organization = Organization.objects.create(name=organization_name)

        # 3. Create a user
        new_user = cls(
            account=new_account,
            organization=new_organization,
            email=email,
            full_name=full_name,
            role=cls.Role.OWNER,
        )
        new_user.save()
        return new_user

    """
    Django Admin Required fields
    """

    def has_perm(self, perm, obj=None):
        return True

    def has_module_perms(self, app_label):
        return True

    @property
    def is_staff(self):
        return self.is_superuser
