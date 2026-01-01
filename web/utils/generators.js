import { createAvatar } from '@dicebear/core'
import { notionists } from '@dicebear/collection'

export function generateAvatarDataURI(seed, { size = 96, radius = 0, backgroundColor } = {}) {
  const avatar = createAvatar(notionists, {
    seed,
    size,
    radius,
    ...(backgroundColor && { backgroundColor: [backgroundColor] }),
  })

  return avatar.toDataUri()
}
