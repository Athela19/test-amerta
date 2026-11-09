import { useState } from 'react'

type Props = {
  className?: string
}

export default function BrandLogo({ className }: Props) {
  const [failed, setFailed] = useState(false)
  const classes = className ? `brand ${className}` : 'brand'

  if (failed) {
    return (
      <span className={classes}>
        Amerta<span>Technology</span>
      </span>
    )
  }

  return (
    <img
      className={classes}
      src="/logo.png"
      alt="Amerta Technology"
      onError={() => setFailed(true)}
    />
  )
}
